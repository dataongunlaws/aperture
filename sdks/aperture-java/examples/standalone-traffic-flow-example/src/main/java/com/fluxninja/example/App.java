package com.fluxninja.example;

import com.fluxninja.aperture.sdk.*;
import io.grpc.ConnectivityState;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;
import spark.Spark;

public class App {
    public static final String DEFAULT_APP_PORT = "8080";
    public static final String DEFAULT_AGENT_HOST = "localhost";
    public static final String DEFAULT_AGENT_PORT = "8089";
    public static final String DEFAULT_FEATURE_NAME = "awesome_feature";
    public static final String DEFAULT_INSECURE_GRPC = "true";
    public static final String DEFAULT_ROOT_CERT = "";

    private final ApertureSDK apertureSDK;
    private final ManagedChannel channel;
    private final String controlPointName;

    public App(ApertureSDK apertureSDK, ManagedChannel channel, String controlPointName) {
        this.apertureSDK = apertureSDK;
        this.channel = channel;
        this.controlPointName = controlPointName;
    }

    public static void main(String[] args) {
        String agentHost = System.getenv("FN_AGENT_HOST");
        if (agentHost == null) {
            agentHost = DEFAULT_AGENT_HOST;
        }
        String agentPort = System.getenv("FN_AGENT_PORT");
        if (agentPort == null) {
            agentPort = DEFAULT_AGENT_PORT;
        }
        String insecureGrpcString = System.getenv("FN_INSECURE_GRPC");
        if (insecureGrpcString == null) {
            insecureGrpcString = DEFAULT_INSECURE_GRPC;
        }
        boolean insecureGrpc = Boolean.parseBoolean(insecureGrpcString);

        String rootCertFile = System.getenv("FN_ROOT_CERTIFICATE_FILE");
        if (rootCertFile == null) {
            rootCertFile = DEFAULT_ROOT_CERT;
        }

        String target = String.format("%s:%s", agentHost, agentPort);
        final ManagedChannel channel = ManagedChannelBuilder.forTarget(target).build();

        ApertureSDK apertureSDK;
        try {
            apertureSDK =
                    ApertureSDK.builder()
                            .setHost(agentHost)
                            .setPort(Integer.parseInt(agentPort))
                            .setFlowTimeout(Duration.ofMillis(1000))
                            .useInsecureGrpc(insecureGrpc)
                            .setRootCertificateFile(rootCertFile)
                            .build();
        } catch (ApertureSDKException e) {
            e.printStackTrace();
            return;
        }

        String featureName = System.getenv("FN_FEATURE_NAME");
        if (featureName == null) {
            featureName = DEFAULT_FEATURE_NAME;
        }

        App app = new App(apertureSDK, channel, featureName);
        String appPort = System.getenv("FN_APP_PORT");
        if (appPort == null) {
            appPort = DEFAULT_APP_PORT;
        }
        Spark.port(Integer.parseInt(appPort));
        Spark.get("/super", app::handleSuperAPI);
        Spark.get("/connected", app::handleConnectedAPI);
        Spark.get("/health", app::handleHealthAPI);
    }

    private String handleSuperAPI(spark.Request req, spark.Response res) {
        Map<String, String> labels = new HashMap<>();

        // do some business logic to collect labels
        labels.put("user", "kenobi");

        Map<String, String> allHeaders = new HashMap<>();
        for (String headerName : req.headers()) {
            allHeaders.put(headerName, req.headers(headerName));
        }
        allHeaders.putAll(labels);

        TrafficFlowRequestBuilder trafficFlowRequestBuilder = TrafficFlowRequest.newBuilder();

        trafficFlowRequestBuilder
                .setControlPoint(this.controlPointName)
                .setHttpMethod(req.requestMethod())
                .setHttpHost(req.host())
                .setHttpProtocol(req.protocol())
                .setHttpPath(req.pathInfo())
                .setHttpScheme(req.scheme())
                .setHttpSize(req.contentLength())
                .setHttpHeaders(allHeaders)
                .setSource(req.ip(), req.port(), "TCP")
                .setDestination(req.raw().getLocalAddr(), req.raw().getLocalPort(), "TCP");

        TrafficFlowRequest apertureRequest = trafficFlowRequestBuilder.build();

        // StartFlow performs a flowcontrolv1.CheckHTTP call to Aperture Agent. It returns a
        // TrafficFlow.
        TrafficFlow flow = this.apertureSDK.startTrafficFlow(apertureRequest);

        // See whether flow was accepted by Aperture Agent.
        if (flow.shouldRun()) {
            // Simulate work being done
            try {
                res.status(202);
                Thread.sleep(2000);
                // Need to call end() on the Flow in order to provide telemetry to Aperture
                // Agent for completing the control loop.
                // The first argument captures whether the feature captured by the Flow was
                // successful or resulted in an error.
                // The second argument is error message for further diagnosis.
                flow.end(FlowStatus.OK);
            } catch (InterruptedException | ApertureSDKException e) {
                e.printStackTrace();
            }
        } else {
            // Flow has been rejected by Aperture Agent.
            try {
                res.status(flow.getRejectionHttpStatusCode());
                flow.end(FlowStatus.Error);
            } catch (ApertureSDKException e) {
                e.printStackTrace();
            }
        }
        return "";
    }

    private String handleConnectedAPI(spark.Request req, spark.Response res) {
        ConnectivityState state = this.channel.getState(true);
        return state.toString();
    }

    private String handleHealthAPI(spark.Request req, spark.Response res) {
        return "Healthy";
    }
}
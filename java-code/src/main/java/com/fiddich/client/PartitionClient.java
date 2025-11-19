package com.fiddich.client;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fiddich.model.*;
import com.fiddich.model.dto.request.PartitionRequest;
import com.fiddich.model.dto.resonse.PartitionResponse;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.List;

import static com.fiddich.util.HttpRequester.send;
import static com.fiddich.util.JsonUtil.deserialize;
import static com.fiddich.util.JsonUtil.serialize;
import static com.fiddich.client.ApiPaths.*;

public class PartitionClient {

    public static ResponseFormat<List<PartitionResponse>> partitionGoods(PartitionRequest requestBody) {
        String jsonBody = serialize(requestBody);

        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(createPath(GOODS_PARTITION)))
                .POST(HttpRequest.BodyPublishers.ofString(jsonBody))
                .build();

        HttpResponse<String> response = send(request);

        return deserialize(
                response.body(),
                new TypeReference<ResponseFormat<List<PartitionResponse>>>() {}
        );
    }
}

package com.fiddich.client;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fiddich.model.dto.request.PaymentRequest;
import com.fiddich.util.JsonUtil;
import com.fiddich.model.*;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import static com.fiddich.util.HttpRequester.send;
import static com.fiddich.util.JsonUtil.deserialize;
import static com.fiddich.client.ApiPaths.*;

public class PaymentClient {

    public static ResponseFormat<Void> paymentGoods(PaymentRequest paymentRequest) {
        String jsonBody = JsonUtil.serialize(paymentRequest);

        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(createPath(PAYMENT)))
                .POST(HttpRequest.BodyPublishers.ofString(jsonBody))
                .build();

        HttpResponse<String> response = send(request);

        System.out.println(paymentRequest.goodsList().stream()
                .map(Goods::getName)
                .toList());
        System.out.println();

        return deserialize(
                response.body(),
                new TypeReference<ResponseFormat<Void>>() {}
        );
    }
}

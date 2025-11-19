package com.fiddich.client;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fiddich.model.dto.resonse.DiscountInfoResponse;
import com.fiddich.model.ResponseFormat;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.List;

import static com.fiddich.util.HttpRequester.send;
import static com.fiddich.util.JsonUtil.deserialize;
import static com.fiddich.client.ApiPaths.*;

public class DiscountPolicyClient {

    public static ResponseFormat<List<DiscountInfoResponse>> getByCoupon() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(createPath(COUPON_DISCOUNT_POLICY)))
                .GET()
                .build();

        HttpResponse<String> response = send(request);

        return deserialize(
                response.body(),
                new TypeReference<ResponseFormat<List<DiscountInfoResponse>>>() {}
        );
    }

    public static ResponseFormat<List<DiscountInfoResponse>> getByCard() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(createPath(CARD_DISCOUNT_POLICY)))
                .GET()
                .build();

        HttpResponse<String> response = send(request);

        return deserialize(
                response.body(),
                new TypeReference<ResponseFormat<List<DiscountInfoResponse>>>() {}
        );
    }
}

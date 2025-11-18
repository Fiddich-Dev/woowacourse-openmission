package com.fiddich.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fiddich.model.DiscountInfoResponse;
import com.fiddich.model.ResponseFormat;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.List;

public class DiscountPolicy {

    private final HttpClient client = HttpClient.newHttpClient();
    private final ObjectMapper mapper = new ObjectMapper();

    public ResponseFormat<List<DiscountInfoResponse>> byCoupon() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:8080/coupon-discount-policy"))
                .GET()
                .build();

        HttpResponse<String> response = sendRequest(request);

        return parseJson(
                response.body(),
                new TypeReference<ResponseFormat<List<DiscountInfoResponse>>>() {}
        );
    }

    public ResponseFormat<List<DiscountInfoResponse>> byCard() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:8080/card-discount-policy"))
                .GET()
                .build();

        HttpResponse<String> response = sendRequest(request);

        return parseJson(
                response.body(),
                new TypeReference<ResponseFormat<List<DiscountInfoResponse>>>() {}
        );
    }

    private HttpResponse<String> sendRequest(HttpRequest request) {
        try {
            return client.send(request, HttpResponse.BodyHandlers.ofString());
        } catch (IOException e) {
            throw new RuntimeException(e);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }

    private <T> T parseJson(String body, TypeReference<T> typeReference) {
        try {
            return mapper.readValue(body, typeReference);
        } catch (JsonProcessingException e) {
            throw new RuntimeException(e);
        }
    }
}

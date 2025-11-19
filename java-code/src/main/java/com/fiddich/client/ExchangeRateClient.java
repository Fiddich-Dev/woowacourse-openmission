package com.fiddich.client;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fiddich.model.dto.resonse.ExchangeRateResponse;
import com.fiddich.model.ResponseFormat;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import static com.fiddich.util.HttpRequester.send;
import static com.fiddich.util.JsonUtil.deserialize;
import static com.fiddich.client.ApiPaths.*;

public class ExchangeRateClient {

    public static ResponseFormat<ExchangeRateResponse> getUsdExchangeRate() {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(createPath(EXCHANGE_RATE)))
                .GET()
                .build();

        HttpResponse<String> response = send(request);

        return deserialize(
                response.body(),
                new TypeReference<ResponseFormat<ExchangeRateResponse>>() {}
        );
    }
}

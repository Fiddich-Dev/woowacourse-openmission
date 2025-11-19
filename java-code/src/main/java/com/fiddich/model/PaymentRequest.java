package com.fiddich.model;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record PaymentRequest(
        List<Goods> goods,
        @JsonProperty("payment_amount") String paymentAmount
) {

}

package com.fiddich.model.dto.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fiddich.model.Goods;

import java.util.List;

public record PaymentRequest(
        List<Goods> goods,
        @JsonProperty("payment_amount") String paymentAmount
) {

}

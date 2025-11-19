package com.fiddich.model.dto.resonse;

import com.fasterxml.jackson.annotation.JsonProperty;

public record DiscountInfoResponse(
        @JsonProperty("threshold") String threshold,
        @JsonProperty("discount_amount") String discountAmount
) {}

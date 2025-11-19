package com.fiddich.model;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record PartitionResponse(
        List<Goods> goods,
        @JsonProperty("before_price") String beforePrice,
        @JsonProperty("after_price") String afterPrice) {
}

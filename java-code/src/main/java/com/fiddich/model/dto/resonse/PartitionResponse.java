package com.fiddich.model.dto.resonse;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fiddich.model.Goods;

import java.util.List;

public record PartitionResponse(
        @JsonProperty("goods_list") List<Goods> goodsList,
        @JsonProperty("before_price") String beforePrice,
        @JsonProperty("after_price") String afterPrice) {
}

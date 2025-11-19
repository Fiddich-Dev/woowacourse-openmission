package com.fiddich.model.dto.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fiddich.model.Goods;

import java.util.List;

public record PartitionRequest(
        @JsonProperty("goods_list") List<Goods> goodsList
) {}

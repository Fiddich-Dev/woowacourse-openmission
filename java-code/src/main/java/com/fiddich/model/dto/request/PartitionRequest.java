package com.fiddich.model.dto.request;

import com.fiddich.model.Goods;

import java.util.List;

public record PartitionRequest(List<Goods> goods) {}

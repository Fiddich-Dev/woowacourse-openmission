package com.fiddich.model;

import java.math.BigDecimal;

@FunctionalInterface
public interface Discount {
    BigDecimal apply(BigDecimal price);
}

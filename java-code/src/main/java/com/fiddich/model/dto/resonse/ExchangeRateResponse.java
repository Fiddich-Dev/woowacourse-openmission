package com.fiddich.model.dto.resonse;

import com.fasterxml.jackson.annotation.JsonProperty;

public record ExchangeRateResponse(
        @JsonProperty("result") Integer result,
        @JsonProperty("cur_unit") String currencyUnit,
        @JsonProperty("ttb") String transferBuying,
        @JsonProperty("tts") String transferSelling,
        @JsonProperty("cur_nm") String currencyName
) {}

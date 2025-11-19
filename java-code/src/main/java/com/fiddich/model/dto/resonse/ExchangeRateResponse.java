package com.fiddich.model.dto.resonse;

import com.fasterxml.jackson.annotation.JsonProperty;

public record ExchangeRateResponse(
        @JsonProperty("result") Integer result,
        @JsonProperty("cur_unit") String currencyUnit,
        @JsonProperty("ttb") String transferBuying,
        @JsonProperty("tts") String transferSelling,
        @JsonProperty("cur_nm") String currencyName
) {

}

/*
        "result": 1,
        "cur_unit": "USD",
        "ttb": "1,450.54",
        "tts": "1,479.85",
        "cur_nm": "미국 달러"
 */

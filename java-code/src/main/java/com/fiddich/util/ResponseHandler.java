package com.fiddich.util;

import com.fiddich.model.ResponseFormat;

public class ResponseHandler {

    public static <T> T handle(ResponseFormat<T> response) {
        int status = response.getStatus();

        if (status >= 400 && status <= 499) {
            throw new IllegalArgumentException(response.getMessage());
        }
        if (status >= 500 && status <= 599) {
            throw new IllegalStateException(response.getMessage());
        }

        return response.getContent();
    }
}

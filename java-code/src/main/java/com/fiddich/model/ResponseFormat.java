package com.fiddich.model;

public class ResponseFormat<T> {
    private Integer status;
    private String message;
    private T content;

    public ResponseFormat() {}

    public Integer getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }

    public T getContent() {
        return content;
    }
}


package com.fiddich.client;

import com.fiddich.model.dto.request.PaymentRequest;
import com.fiddich.model.dto.resonse.PartitionResponse;
import com.fiddich.util.ResponseHandler;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;

public class PaymentExecutor {
    private static final String PAYMENT_FAIL_MESSAGE = "결제 실패: ";
    private static final String INTERRUPTED_MESSAGE = "작업 중단";

    private final ExecutorService executorService;

    public PaymentExecutor(int threadCount) {
        this.executorService = Executors.newFixedThreadPool(threadCount);
    }

    public List<PartitionResponse> executePayments(List<PartitionResponse> partitions) {
        List<Future<PartitionResponse>> futures = new ArrayList<>();

        for (PartitionResponse p : partitions) {
            futures.add(executorService.submit(createTask(p)));
        }

        List<PartitionResponse> successResults = new ArrayList<>();

        for (Future<PartitionResponse> future : futures) {
            try {
                successResults.add(future.get());
            } catch (ExecutionException e) {
                System.out.println(PAYMENT_FAIL_MESSAGE + e.getCause().getMessage());
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                System.out.println(INTERRUPTED_MESSAGE);
            }
        }

        shutdown();
        return successResults;
    }

    private Callable<PartitionResponse> createTask(PartitionResponse partitionResponse) {
        return () -> {
            ResponseHandler.voidHandle(
                    PaymentClient.paymentGoods(
                            new PaymentRequest(partitionResponse.goodsList(), partitionResponse.afterPrice())
                    )
            );
            return partitionResponse;
        };
    }

    private void shutdown() {
        executorService.shutdown();
    }
}

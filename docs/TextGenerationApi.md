# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TextGeneration**](TextGenerationApi.md#TextGeneration) | **Post** /api/v1/text-generation | Text Generation - Multi-Provider

# **TextGeneration**
> TextGenerationResponse TextGeneration(ctx, provider, model, input, topK, topP, temperature, repetitionPenalty, dtype, maxLength, maxNewTokens, minLength, minNewTokens, doSample, numBeams, noRepeatNgramSize, contextWindowSize, slidingWindowSize, attentionSinkSize, frequencyPenalty, presencePenalty, bosTokenId, maxTokens, randomSeed)
Text Generation - Multi-Provider

Generate text using multiple AI providers (Transformers.js, WebLLM, MediaPipe). Use the 'provider' field to specify which AI provider to use for text generation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 
  **model** | **string**|  | 
  **input** | **string**|  | 
  **topK** | **int32**|  | 
  **topP** | **float64**|  | 
  **temperature** | **float64**|  | 
  **repetitionPenalty** | **float64**|  | 
  **dtype** | **string**|  | 
  **maxLength** | **int32**|  | 
  **maxNewTokens** | **int32**|  | 
  **minLength** | **int32**|  | 
  **minNewTokens** | **int32**|  | 
  **doSample** | **bool**|  | 
  **numBeams** | **int32**|  | 
  **noRepeatNgramSize** | **int32**|  | 
  **contextWindowSize** | **int32**|  | 
  **slidingWindowSize** | **int32**|  | 
  **attentionSinkSize** | **int32**|  | 
  **frequencyPenalty** | **float64**|  | 
  **presencePenalty** | **float64**|  | 
  **bosTokenId** | **int32**|  | 
  **maxTokens** | **int32**|  | 
  **randomSeed** | **int32**|  | 

### Return type

[**TextGenerationResponse**](TextGenerationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


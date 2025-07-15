# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TextToSpeech**](TextToSpeechApi.md#TextToSpeech) | **Post** /api/v1/text-to-speech | Text-to-Speech

# **TextToSpeech**
> []TtsResponse TextToSpeech(ctx, model, dtype, input, voice, stream)
Text-to-Speech

Generate natural speech from text using MMS or Kokoro models. Supports multiple languages and premium voices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**|  | 
  **dtype** | **string**|  | 
  **input** | **string**|  | 
  **voice** | **string**|  | 
  **stream** | **bool**|  | 

### Return type

[**[]TtsResponse**](TTSResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Translation**](TranslationApi.md#Translation) | **Post** /api/v1/translation | Translation

# **Translation**
> TranslationResponse Translation(ctx, model, dtype, input, srcLang, tgtLang)
Translation

Translate text between 200+ languages using NLLB models. Uses FLORES200 format for language codes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**|  | 
  **dtype** | **string**|  | 
  **input** | **string**|  | 
  **srcLang** | **string**|  | 
  **tgtLang** | **string**|  | 

### Return type

[**TranslationResponse**](TranslationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


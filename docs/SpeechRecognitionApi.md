# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpeechToText**](SpeechRecognitionApi.md#SpeechToText) | **Post** /api/v1/speech-recognition | Speech Recognition (Speech-to-Text)

# **SpeechToText**
> []SttChunk SpeechToText(ctx, model, dtype, input, returnTimestamps, stream, chunkLengthS, strideLengthS, forceFullSequences, language, task, numFrames)
Speech Recognition (Speech-to-Text)

Convert audio files to text using Whisper models. Supports MP3, WAV, M4A and other audio formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**|  | 
  **dtype** | **string**|  | 
  **input** | [**Object**](.md)|  | 
  **returnTimestamps** | **string**|  | 
  **stream** | **bool**|  | 
  **chunkLengthS** | **int32**|  | 
  **strideLengthS** | **int32**|  | 
  **forceFullSequences** | **bool**|  | 
  **language** | **string**|  | 
  **task** | **string**|  | 
  **numFrames** | **int32**|  | 

### Return type

[**[]SttChunk**](STTChunk.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


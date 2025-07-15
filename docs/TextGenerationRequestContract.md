# TextGenerationRequestContract

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The AI provider to use | [default to null]
**Model** | **string** | The AI model to use for processing | [default to null]
**Input** | **string** | Input text or messages for generation | [default to null]
**TopK** | **int32** | The number of highest probability vocabulary tokens to keep for top-k-filtering | [optional] [default to null]
**TopP** | **float64** | If set to float &lt; 1, only the smallest set of most probable tokens with probabilities that add up to top_p or higher are kept for generation | [optional] [default to null]
**Temperature** | **float64** | The value used to modulate the next token probabilities | [optional] [default to null]
**RepetitionPenalty** | **float64** | Parameter for repetition penalty. 1.0 means no penalty | [optional] [default to null]
**Dtype** | **string** | Quantization level (e.g., &#x27;fp16&#x27;, &#x27;q4&#x27;, &#x27;q8&#x27;) - Transformers only | [optional] [default to null]
**MaxLength** | **int32** | Maximum length the generated tokens can have - Transformers only | [optional] [default to null]
**MaxNewTokens** | **int32** | Maximum number of tokens to generate - Transformers only | [optional] [default to null]
**MinLength** | **int32** | Minimum length of the sequence to be generated - Transformers only | [optional] [default to null]
**MinNewTokens** | **int32** | Minimum numbers of tokens to generate - Transformers only | [optional] [default to null]
**DoSample** | **bool** | Whether to use sampling - Transformers only | [optional] [default to null]
**NumBeams** | **int32** | Number of beams for beam search - Transformers only | [optional] [default to null]
**NoRepeatNgramSize** | **int32** | If &gt; 0, all ngrams of that size can only occur once - Transformers only | [optional] [default to null]
**ContextWindowSize** | **int32** | Size of the context window for the model - WebLLM only | [optional] [default to null]
**SlidingWindowSize** | **int32** | Size of the sliding window for attention - WebLLM only | [optional] [default to null]
**AttentionSinkSize** | **int32** | Size of the attention sink - WebLLM only | [optional] [default to null]
**FrequencyPenalty** | **float64** | Penalty for token frequency - WebLLM only | [optional] [default to null]
**PresencePenalty** | **float64** | Penalty for token presence - WebLLM only | [optional] [default to null]
**BosTokenId** | **int32** | Beginning of sequence token ID - WebLLM only | [optional] [default to null]
**MaxTokens** | **int32** | Maximum number of tokens to generate - MediaPipe only | [optional] [default to null]
**RandomSeed** | **int32** | Random seed for reproducible results - MediaPipe only | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


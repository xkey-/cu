package main

var ignored = map[string]struct{}{
	"cudnnGetVersion":              {},
	"cudnnGetCudartVersion":        {},
	"cudnnGetErrorString":          {},
	"cudnnQueryRuntimeError":       {},
	"cudnnGetProperty":             {},
	"cudnnCreate":                  {},
	"cudnnDestroy":                 {},
	"cudnnSetStream":               {},
	"cudnnGetStream":               {},
	"cudnnCreateTensorDescriptor":  {},
	"cudnnSetTensor4dDescriptor":   {},
	"cudnnSetTensor4dDescriptorEx": {},
	"cudnnGetTensor4dDescriptor":   {},
	"cudnnSetTensorNdDescriptor":   {},
	"cudnnSetTensorNdDescriptorEx": {},
	"cudnnGetTensorNdDescriptor":   {},
	"cudnnGetTensorSizeInBytes":    {},
	"cudnnDestroyTensorDescriptor": {},
	"cudnnTransformTensor":         {},
	// "cudnnAddTensor":                                     {},
	"cudnnCreateOpTensorDescriptor":  {},
	"cudnnSetOpTensorDescriptor":     {},
	"cudnnGetOpTensorDescriptor":     {},
	"cudnnDestroyOpTensorDescriptor": {},
	// "cudnnOpTensor":                                      {},
	"cudnnCreateReduceTensorDescriptor": {},
	"cudnnSetReduceTensorDescriptor":    {},
	// "cudnnGetReduceTensorDescriptor":                     {},
	// "cudnnDestroyReduceTensorDescriptor":                 {},
	// "cudnnGetReductionIndicesSize":                       {},
	// "cudnnGetReductionWorkspaceSize":                     {},
	// "cudnnReduceTensor":                                  {},
	"cudnnSetTensor": {},
	// "cudnnScaleTensor":                                   {},
	"cudnnCreateFilterDescriptor":                 {},
	"cudnnSetFilter4dDescriptor":                  {},
	"cudnnGetFilter4dDescriptor":                  {},
	"cudnnSetFilterNdDescriptor":                  {},
	"cudnnGetFilterNdDescriptor":                  {},
	"cudnnDestroyFilterDescriptor":                {},
	"cudnnCreateConvolutionDescriptor":            {},
	"cudnnSetConvolutionMathType":                 {},
	"cudnnGetConvolutionMathType":                 {},
	"cudnnSetConvolutionGroupCount":               {},
	"cudnnGetConvolutionGroupCount":               {},
	"cudnnSetConvolution2dDescriptor":             {},
	"cudnnGetConvolution2dDescriptor":             {},
	"cudnnGetConvolution2dForwardOutputDim":       {},
	"cudnnSetConvolutionNdDescriptor":             {},
	"cudnnGetConvolutionNdDescriptor":             {},
	"cudnnGetConvolutionNdForwardOutputDim":       {},
	"cudnnDestroyConvolutionDescriptor":           {},
	"cudnnGetConvolutionForwardAlgorithmMaxCount": {},
	// "cudnnFindConvolutionForwardAlgorithm":               {},
	// "cudnnFindConvolutionForwardAlgorithmEx":             {},
	// "cudnnGetConvolutionForwardAlgorithm":                {},
	// "cudnnGetConvolutionForwardAlgorithm_v7":             {},
	// "cudnnGetConvolutionForwardWorkspaceSize":            {},
	// "cudnnConvolutionForward":                            {},
	// "cudnnConvolutionBiasActivationForward":              {},
	// "cudnnConvolutionBackwardBias":                       {},
	// "cudnnGetConvolutionBackwardFilterAlgorithmMaxCount": {},
	// "cudnnFindConvolutionBackwardFilterAlgorithm":        {},
	// "cudnnFindConvolutionBackwardFilterAlgorithmEx":      {},
	// "cudnnGetConvolutionBackwardFilterAlgorithm":         {},
	// "cudnnGetConvolutionBackwardFilterAlgorithm_v7":      {},
	// "cudnnGetConvolutionBackwardFilterWorkspaceSize":     {},
	// "cudnnConvolutionBackwardFilter":                     {},
	// "cudnnGetConvolutionBackwardDataAlgorithmMaxCount":   {},
	// "cudnnFindConvolutionBackwardDataAlgorithm":          {},
	// "cudnnFindConvolutionBackwardDataAlgorithmEx":        {},
	// "cudnnGetConvolutionBackwardDataAlgorithm":           {},
	// "cudnnGetConvolutionBackwardDataAlgorithm_v7":        {},
	// "cudnnGetConvolutionBackwardDataWorkspaceSize":       {},
	// "cudnnConvolutionBackwardData":                       {},
	// "cudnnIm2Col":                                        {},
	// "cudnnSoftmaxForward":                                {},
	// "cudnnSoftmaxBackward":                               {},
	// "cudnnCreatePoolingDescriptor":                       {},
	// "cudnnSetPooling2dDescriptor": {},
	// "cudnnGetPooling2dDescriptor": {},
	// "cudnnSetPoolingNdDescriptor": {},
	// "cudnnGetPoolingNdDescriptor": {},
	// "cudnnGetPoolingNdForwardOutputDim":                  {},
	// "cudnnGetPooling2dForwardOutputDim":                  {},
	// "cudnnDestroyPoolingDescriptor":                      {},
	// "cudnnPoolingForward":                                {},
	// "cudnnPoolingBackward":                               {},
	// "cudnnCreateActivationDescriptor":                    {},
	// "cudnnSetActivationDescriptor": {},
	// "cudnnGetActivationDescriptor": {},
	// "cudnnDestroyActivationDescriptor":                   {},
	// "cudnnActivationForward":                             {},
	// "cudnnActivationBackward":                            {},
	// "cudnnCreateLRNDescriptor":                           {},
	// "cudnnSetLRNDescriptor": {},
	// "cudnnGetLRNDescriptor": {},
	// "cudnnDestroyLRNDescriptor":                          {},
	// "cudnnLRNCrossChannelForward":                        {},
	// "cudnnLRNCrossChannelBackward":                       {},
	// "cudnnDivisiveNormalizationForward":                  {},
	// "cudnnDivisiveNormalizationBackward":                 {},
	// "cudnnDeriveBNTensorDescriptor":                      {},
	// "cudnnBatchNormalizationForwardTraining":             {},
	// "cudnnBatchNormalizationForwardInference":            {},
	// "cudnnBatchNormalizationBackward":                    {},
	// "cudnnCreateSpatialTransformerDescriptor":            {},
	// "cudnnSetSpatialTransformerNdDescriptor": {},
	// "cudnnDestroySpatialTransformerDescriptor":           {},
	// "cudnnSpatialTfGridGeneratorForward":                 {},
	// "cudnnSpatialTfGridGeneratorBackward":                {},
	// "cudnnSpatialTfSamplerForward":                       {},
	// "cudnnSpatialTfSamplerBackward":                      {},
	// "cudnnCreateDropoutDescriptor":                       {},
	// "cudnnDestroyDropoutDescriptor":                      {},
	// "cudnnDropoutGetStatesSize":                          {},
	// "cudnnDropoutGetReserveSpaceSize":                    {},
	// "cudnnSetDropoutDescriptor": {},
	// "cudnnRestoreDropoutDescriptor":                      {},
	// "cudnnGetDropoutDescriptor":                          {},
	// "cudnnDropoutForward":                                {},
	// "cudnnDropoutBackward":                               {},
	// "cudnnCreateRNNDescriptor":                           {},
	// "cudnnDestroyRNNDescriptor":                          {},
	// "cudnnCreatePersistentRNNPlan":                       {},
	// "cudnnSetPersistentRNNPlan": {},
	// "cudnnDestroyPersistentRNNPlan":                      {},
	// "cudnnSetRNNDescriptor": {},
	// "cudnnGetRNNDescriptor":                              {},
	// "cudnnSetRNNMatrixMathType": {},
	// "cudnnGetRNNWorkspaceSize":                           {},
	// "cudnnGetRNNTrainingReserveSize":                     {},
	// "cudnnGetRNNParamsSize":                              {},
	// "cudnnGetRNNLinLayerMatrixParams":                    {},
	// "cudnnGetRNNLinLayerBiasParams":                      {},
	// "cudnnRNNForwardInference":                           {},
	// "cudnnRNNForwardTraining":                            {},
	// "cudnnRNNBackwardData":                               {},
	// "cudnnRNNBackwardWeights":                            {},
	// "cudnnCreateCTCLossDescriptor":                       {},
	// "cudnnSetCTCLossDescriptor": {},
	// "cudnnGetCTCLossDescriptor":                          {},
	// "cudnnDestroyCTCLossDescriptor":                      {},
	// "cudnnCTCLoss":                                       {},
	// "cudnnGetCTCLossWorkspaceSize":                       {},
	// "cudnnSetRNNDescriptor_v6": {},
	// "cudnnSetRNNDescriptor_v5": {},
}

func init() {

	fnNameMap = map[string]string{
		"cudnnGetVersion":                                    "GetVersion",
		"cudnnGetCudartVersion":                              "GetCudartVersion",
		"cudnnGetErrorString":                                "GetErrorString",
		"cudnnQueryRuntimeError":                             "QueryRuntimeError",
		"cudnnGetProperty":                                   "GetProperty",
		"cudnnCreate":                                        "Create",
		"cudnnDestroy":                                       "Destroy",
		"cudnnSetStream":                                     "SetStream",
		"cudnnGetStream":                                     "GetStream",
		"cudnnCreateTensorDescriptor":                        "CreateTensorDescriptor",
		"cudnnSetTensor4dDescriptor":                         "SetTensor4dDescriptor",
		"cudnnSetTensor4dDescriptorEx":                       "SetTensor4dDescriptorEx",
		"cudnnGetTensor4dDescriptor":                         "GetTensor4dDescriptor",
		"cudnnSetTensorNdDescriptor":                         "SetTensorNdDescriptor",
		"cudnnSetTensorNdDescriptorEx":                       "SetTensorNdDescriptorEx",
		"cudnnGetTensorNdDescriptor":                         "GetTensorNdDescriptor",
		"cudnnGetTensorSizeInBytes":                          "GetTensorSizeInBytes",
		"cudnnDestroyTensorDescriptor":                       "DestroyTensorDescriptor",
		"cudnnTransformTensor":                               "TransformTensor",
		"cudnnAddTensor":                                     "AddTensor",
		"cudnnCreateOpTensorDescriptor":                      "CreateOpTensorDescriptor",
		"cudnnSetOpTensorDescriptor":                         "SetOpTensorDescriptor",
		"cudnnGetOpTensorDescriptor":                         "GetOpTensorDescriptor",
		"cudnnDestroyOpTensorDescriptor":                     "DestroyOpTensorDescriptor",
		"cudnnOpTensor":                                      "OpTensor",
		"cudnnCreateReduceTensorDescriptor":                  "CreateReduceTensorDescriptor",
		"cudnnSetReduceTensorDescriptor":                     "SetReduceTensorDescriptor",
		"cudnnGetReduceTensorDescriptor":                     "GetReduceTensorDescriptor",
		"cudnnDestroyReduceTensorDescriptor":                 "DestroyReduceTensorDescriptor",
		"cudnnGetReductionIndicesSize":                       "GetReductionIndicesSize",
		"cudnnGetReductionWorkspaceSize":                     "GetReductionWorkspaceSize",
		"cudnnReduceTensor":                                  "ReduceTensor",
		"cudnnSetTensor":                                     "SetTensor",
		"cudnnScaleTensor":                                   "ScaleTensor",
		"cudnnCreateFilterDescriptor":                        "CreateFilterDescriptor",
		"cudnnSetFilter4dDescriptor":                         "SetFilter4dDescriptor",
		"cudnnGetFilter4dDescriptor":                         "GetFilter4dDescriptor",
		"cudnnSetFilterNdDescriptor":                         "SetFilterNdDescriptor",
		"cudnnGetFilterNdDescriptor":                         "GetFilterNdDescriptor",
		"cudnnDestroyFilterDescriptor":                       "DestroyFilterDescriptor",
		"cudnnCreateConvolutionDescriptor":                   "CreateConvolutionDescriptor",
		"cudnnSetConvolutionMathType":                        "SetConvolutionMathType",
		"cudnnGetConvolutionMathType":                        "GetConvolutionMathType",
		"cudnnSetConvolutionGroupCount":                      "SetConvolutionGroupCount",
		"cudnnGetConvolutionGroupCount":                      "GetConvolutionGroupCount",
		"cudnnSetConvolution2dDescriptor":                    "SetConvolution2dDescriptor",
		"cudnnGetConvolution2dDescriptor":                    "GetConvolution2dDescriptor",
		"cudnnGetConvolution2dForwardOutputDim":              "GetConvolution2dForwardOutputDim",
		"cudnnSetConvolutionNdDescriptor":                    "SetConvolutionNdDescriptor",
		"cudnnGetConvolutionNdDescriptor":                    "GetConvolutionNdDescriptor",
		"cudnnGetConvolutionNdForwardOutputDim":              "GetConvolutionNdForwardOutputDim",
		"cudnnDestroyConvolutionDescriptor":                  "DestroyConvolutionDescriptor",
		"cudnnGetConvolutionForwardAlgorithmMaxCount":        "GetConvolutionForwardAlgorithmMaxCount",
		"cudnnFindConvolutionForwardAlgorithm":               "FindConvolutionForwardAlgorithm",
		"cudnnFindConvolutionForwardAlgorithmEx":             "FindConvolutionForwardAlgorithmEx",
		"cudnnGetConvolutionForwardAlgorithm":                "GetConvolutionForwardAlgorithm",
		"cudnnGetConvolutionForwardAlgorithm_v7":             "GetConvolutionForwardAlgorithm_v7",
		"cudnnGetConvolutionForwardWorkspaceSize":            "GetConvolutionForwardWorkspaceSize",
		"cudnnConvolutionForward":                            "ConvolutionForward",
		"cudnnConvolutionBiasActivationForward":              "ConvolutionBiasActivationForward",
		"cudnnConvolutionBackwardBias":                       "ConvolutionBackwardBias",
		"cudnnGetConvolutionBackwardFilterAlgorithmMaxCount": "GetConvolutionBackwardFilterAlgorithmMaxCount",
		"cudnnFindConvolutionBackwardFilterAlgorithm":        "FindConvolutionBackwardFilterAlgorithm",
		"cudnnFindConvolutionBackwardFilterAlgorithmEx":      "FindConvolutionBackwardFilterAlgorithmEx",
		"cudnnGetConvolutionBackwardFilterAlgorithm":         "GetConvolutionBackwardFilterAlgorithm",
		"cudnnGetConvolutionBackwardFilterAlgorithm_v7":      "GetConvolutionBackwardFilterAlgorithm_v7",
		"cudnnGetConvolutionBackwardFilterWorkspaceSize":     "GetConvolutionBackwardFilterWorkspaceSize",
		"cudnnConvolutionBackwardFilter":                     "ConvolutionBackwardFilter",
		"cudnnGetConvolutionBackwardDataAlgorithmMaxCount":   "GetConvolutionBackwardDataAlgorithmMaxCount",
		"cudnnFindConvolutionBackwardDataAlgorithm":          "FindConvolutionBackwardDataAlgorithm",
		"cudnnFindConvolutionBackwardDataAlgorithmEx":        "FindConvolutionBackwardDataAlgorithmEx",
		"cudnnGetConvolutionBackwardDataAlgorithm":           "GetConvolutionBackwardDataAlgorithm",
		"cudnnGetConvolutionBackwardDataAlgorithm_v7":        "GetConvolutionBackwardDataAlgorithm_v7",
		"cudnnGetConvolutionBackwardDataWorkspaceSize":       "GetConvolutionBackwardDataWorkspaceSize",
		"cudnnConvolutionBackwardData":                       "ConvolutionBackwardData",
		"cudnnIm2Col":                                        "Im2Col",
		"cudnnSoftmaxForward":                                "SoftmaxForward",
		"cudnnSoftmaxBackward":                               "SoftmaxBackward",
		"cudnnCreatePoolingDescriptor":                       "CreatePoolingDescriptor",
		"cudnnSetPooling2dDescriptor":                        "SetPooling2dDescriptor",
		"cudnnGetPooling2dDescriptor":                        "GetPooling2dDescriptor",
		"cudnnSetPoolingNdDescriptor":                        "SetPoolingNdDescriptor",
		"cudnnGetPoolingNdDescriptor":                        "GetPoolingNdDescriptor",
		"cudnnGetPoolingNdForwardOutputDim":                  "GetPoolingNdForwardOutputDim",
		"cudnnGetPooling2dForwardOutputDim":                  "GetPooling2dForwardOutputDim",
		"cudnnDestroyPoolingDescriptor":                      "DestroyPoolingDescriptor",
		"cudnnPoolingForward":                                "PoolingForward",
		"cudnnPoolingBackward":                               "PoolingBackward",
		"cudnnCreateActivationDescriptor":                    "CreateActivationDescriptor",
		"cudnnSetActivationDescriptor":                       "SetActivationDescriptor",
		"cudnnGetActivationDescriptor":                       "GetActivationDescriptor",
		"cudnnDestroyActivationDescriptor":                   "DestroyActivationDescriptor",
		"cudnnActivationForward":                             "ActivationForward",
		"cudnnActivationBackward":                            "ActivationBackward",
		"cudnnCreateLRNDescriptor":                           "CreateLRNDescriptor",
		"cudnnSetLRNDescriptor":                              "SetLRNDescriptor",
		"cudnnGetLRNDescriptor":                              "GetLRNDescriptor",
		"cudnnDestroyLRNDescriptor":                          "DestroyLRNDescriptor",
		"cudnnLRNCrossChannelForward":                        "LRNCrossChannelForward",
		"cudnnLRNCrossChannelBackward":                       "LRNCrossChannelBackward",
		"cudnnDivisiveNormalizationForward":                  "DivisiveNormalizationForward",
		"cudnnDivisiveNormalizationBackward":                 "DivisiveNormalizationBackward",
		"cudnnDeriveBNTensorDescriptor":                      "DeriveBNTensorDescriptor",
		"cudnnBatchNormalizationForwardTraining":             "BatchNormalizationForwardTraining",
		"cudnnBatchNormalizationForwardInference":            "BatchNormalizationForwardInference",
		"cudnnBatchNormalizationBackward":                    "BatchNormalizationBackward",
		"cudnnCreateSpatialTransformerDescriptor":            "CreateSpatialTransformerDescriptor",
		"cudnnSetSpatialTransformerNdDescriptor":             "SetSpatialTransformerNdDescriptor",
		"cudnnDestroySpatialTransformerDescriptor":           "DestroySpatialTransformerDescriptor",
		"cudnnSpatialTfGridGeneratorForward":                 "SpatialTfGridGeneratorForward",
		"cudnnSpatialTfGridGeneratorBackward":                "SpatialTfGridGeneratorBackward",
		"cudnnSpatialTfSamplerForward":                       "SpatialTfSamplerForward",
		"cudnnSpatialTfSamplerBackward":                      "SpatialTfSamplerBackward",
		"cudnnCreateDropoutDescriptor":                       "CreateDropoutDescriptor",
		"cudnnDestroyDropoutDescriptor":                      "DestroyDropoutDescriptor",
		"cudnnDropoutGetStatesSize":                          "DropoutGetStatesSize",
		"cudnnDropoutGetReserveSpaceSize":                    "DropoutGetReserveSpaceSize",
		"cudnnSetDropoutDescriptor":                          "SetDropoutDescriptor",
		"cudnnRestoreDropoutDescriptor":                      "RestoreDropoutDescriptor",
		"cudnnGetDropoutDescriptor":                          "GetDropoutDescriptor",
		"cudnnDropoutForward":                                "DropoutForward",
		"cudnnDropoutBackward":                               "DropoutBackward",
		"cudnnCreateRNNDescriptor":                           "CreateRNNDescriptor",
		"cudnnDestroyRNNDescriptor":                          "DestroyRNNDescriptor",
		"cudnnCreatePersistentRNNPlan":                       "CreatePersistentRNNPlan",
		"cudnnSetPersistentRNNPlan":                          "SetPersistentRNNPlan",
		"cudnnDestroyPersistentRNNPlan":                      "DestroyPersistentRNNPlan",
		"cudnnSetRNNDescriptor":                              "SetRNNDescriptor",
		"cudnnGetRNNDescriptor":                              "GetRNNDescriptor",
		"cudnnSetRNNMatrixMathType":                          "SetRNNMatrixMathType",
		"cudnnGetRNNWorkspaceSize":                           "GetRNNWorkspaceSize",
		"cudnnGetRNNTrainingReserveSize":                     "GetRNNTrainingReserveSize",
		"cudnnGetRNNParamsSize":                              "GetRNNParamsSize",
		"cudnnGetRNNLinLayerMatrixParams":                    "GetRNNLinLayerMatrixParams",
		"cudnnGetRNNLinLayerBiasParams":                      "GetRNNLinLayerBiasParams",
		"cudnnRNNForwardInference":                           "RNNForwardInference",
		"cudnnRNNForwardTraining":                            "RNNForwardTraining",
		"cudnnRNNBackwardData":                               "RNNBackwardData",
		"cudnnRNNBackwardWeights":                            "RNNBackwardWeights",
		"cudnnCreateCTCLossDescriptor":                       "CreateCTCLossDescriptor",
		"cudnnSetCTCLossDescriptor":                          "SetCTCLossDescriptor",
		"cudnnGetCTCLossDescriptor":                          "GetCTCLossDescriptor",
		"cudnnDestroyCTCLossDescriptor":                      "DestroyCTCLossDescriptor",
		"cudnnCTCLoss":                                       "CTCLoss",
		"cudnnGetCTCLossWorkspaceSize":                       "GetCTCLossWorkspaceSize",
		"cudnnSetRNNDescriptor_v6":                           "SetRNNDescriptor_v6",
		"cudnnSetRNNDescriptor_v5":                           "SetRNNDescriptor_v5",
	}
	enumMappings = map[string]string{
		"cudnnStatus_t":                         "Status",
		"cudnnErrQueryMode_t":                   "ErrQueryMode",
		"cudnnDataType_t":                       "DataType",
		"cudnnMathType_t":                       "MathType",
		"cudnnNanPropagation_t":                 "NanPropagation",
		"cudnnDeterminism_t":                    "Determinism",
		"cudnnTensorFormat_t":                   "TensorFormat",
		"cudnnOpTensorOp_t":                     "OpTensorOp",
		"cudnnReduceTensorOp_t":                 "ReduceTensorOp",
		"cudnnReduceTensorIndices_t":            "ReduceTensorIndices",
		"cudnnIndicesType_t":                    "IndicesType",
		"cudnnConvolutionMode_t":                "ConvolutionMode",
		"cudnnConvolutionFwdPreference_t":       "ConvolutionFwdPreference",
		"cudnnConvolutionFwdAlgo_t":             "ConvolutionFwdAlgo",
		"cudnnConvolutionBwdFilterPreference_t": "ConvolutionBwdFilterPreference",
		"cudnnConvolutionBwdFilterAlgo_t":       "ConvolutionBwdFilterAlgo",
		"cudnnConvolutionBwdDataPreference_t":   "ConvolutionBwdDataPreference",
		"cudnnConvolutionBwdDataAlgo_t":         "ConvolutionBwdDataAlgo",
		"cudnnSoftmaxAlgorithm_t":               "SoftmaxAlgorithm",
		"cudnnSoftmaxMode_t":                    "SoftmaxMode",
		"cudnnPoolingMode_t":                    "PoolingMode",
		"cudnnActivationMode_t":                 "ActivationMode",
		"cudnnLRNMode_t":                        "LRNMode",
		"cudnnDivNormMode_t":                    "DivNormMode",
		"cudnnBatchNormMode_t":                  "BatchNormMode",
		"cudnnSamplerType_t":                    "SamplerType",
		"cudnnRNNMode_t":                        "RNNMode",
		"cudnnDirectionMode_t":                  "DirectionMode",
		"cudnnRNNInputMode_t":                   "RNNInputMode",
		"cudnnRNNAlgo_t":                        "RNNAlgo",
		"cudnnCTCLossAlgo_t":                    "CTCLossAlgo",
	}
	contextual = map[string]struct{}{
		// "cudnnQueryRuntimeError":                             {},
		// "cudnnCreate":                                        {},
		// "cudnnDestroy":                                       {},
		"cudnnSetStream":                                     {},
		"cudnnGetStream":                                     {},
		"cudnnTransformTensor":                               {},
		"cudnnAddTensor":                                     {},
		"cudnnOpTensor":                                      {},
		"cudnnGetReductionIndicesSize":                       {},
		"cudnnGetReductionWorkspaceSize":                     {},
		"cudnnReduceTensor":                                  {},
		"cudnnSetTensor":                                     {},
		"cudnnScaleTensor":                                   {},
		"cudnnGetConvolutionForwardAlgorithmMaxCount":        {},
		"cudnnFindConvolutionForwardAlgorithm":               {},
		"cudnnFindConvolutionForwardAlgorithmEx":             {},
		"cudnnGetConvolutionForwardAlgorithm":                {},
		"cudnnGetConvolutionForwardAlgorithm_v7":             {},
		"cudnnGetConvolutionForwardWorkspaceSize":            {},
		"cudnnConvolutionForward":                            {},
		"cudnnConvolutionBiasActivationForward":              {},
		"cudnnConvolutionBackwardBias":                       {},
		"cudnnGetConvolutionBackwardFilterAlgorithmMaxCount": {},
		"cudnnFindConvolutionBackwardFilterAlgorithm":        {},
		"cudnnFindConvolutionBackwardFilterAlgorithmEx":      {},
		"cudnnGetConvolutionBackwardFilterAlgorithm":         {},
		"cudnnGetConvolutionBackwardFilterAlgorithm_v7":      {},
		"cudnnGetConvolutionBackwardFilterWorkspaceSize":     {},
		"cudnnConvolutionBackwardFilter":                     {},
		"cudnnGetConvolutionBackwardDataAlgorithmMaxCount":   {},
		"cudnnFindConvolutionBackwardDataAlgorithm":          {},
		"cudnnFindConvolutionBackwardDataAlgorithmEx":        {},
		"cudnnGetConvolutionBackwardDataAlgorithm":           {},
		"cudnnGetConvolutionBackwardDataAlgorithm_v7":        {},
		"cudnnGetConvolutionBackwardDataWorkspaceSize":       {},
		"cudnnConvolutionBackwardData":                       {},
		"cudnnIm2Col":                                        {},
		"cudnnSoftmaxForward":                                {},
		"cudnnSoftmaxBackward":                               {},
		"cudnnPoolingForward":                                {},
		"cudnnPoolingBackward":                               {},
		"cudnnActivationForward":                             {},
		"cudnnActivationBackward":                            {},
		"cudnnLRNCrossChannelForward":                        {},
		"cudnnLRNCrossChannelBackward":                       {},
		"cudnnDivisiveNormalizationForward":                  {},
		"cudnnDivisiveNormalizationBackward":                 {},
		"cudnnBatchNormalizationForwardTraining":             {},
		"cudnnBatchNormalizationForwardInference":            {},
		"cudnnBatchNormalizationBackward":                    {},
		"cudnnSpatialTfGridGeneratorForward":                 {},
		"cudnnSpatialTfGridGeneratorBackward":                {},
		"cudnnSpatialTfSamplerForward":                       {},
		"cudnnSpatialTfSamplerBackward":                      {},
		"cudnnDropoutGetStatesSize":                          {},
		"cudnnSetDropoutDescriptor":                          {},
		"cudnnRestoreDropoutDescriptor":                      {},
		"cudnnGetDropoutDescriptor":                          {},
		"cudnnDropoutForward":                                {},
		"cudnnDropoutBackward":                               {},
		"cudnnSetRNNDescriptor":                              {},
		"cudnnGetRNNWorkspaceSize":                           {},
		"cudnnGetRNNTrainingReserveSize":                     {},
		"cudnnGetRNNParamsSize":                              {},
		"cudnnGetRNNLinLayerMatrixParams":                    {},
		"cudnnGetRNNLinLayerBiasParams":                      {},
		"cudnnRNNForwardInference":                           {},
		"cudnnRNNForwardTraining":                            {},
		"cudnnRNNBackwardData":                               {},
		"cudnnRNNBackwardWeights":                            {},
		"cudnnCTCLoss":                                       {},
		"cudnnGetCTCLossWorkspaceSize":                       {},
		"cudnnSetRNNDescriptor_v6":                           {},
	}

	alphaBetas = map[string]map[int]string{
		"cudnnTransformTensor":                    {4: "beta", 1: "alpha"},
		"cudnnAddTensor":                          {4: "beta", 1: "alpha"},
		"cudnnOpTensor":                           {8: "beta", 5: "alpha2", 2: "alpha1"},
		"cudnnReduceTensor":                       {9: "beta", 6: "alpha"},
		"cudnnScaleTensor":                        {3: "alpha"},
		"cudnnConvolutionForward":                 {10: "beta", 1: "alpha"},
		"cudnnConvolutionBiasActivationForward":   {10: "alpha2", 1: "alpha1"},
		"cudnnConvolutionBackwardBias":            {4: "beta", 1: "alpha"},
		"cudnnConvolutionBackwardFilter":          {10: "beta", 1: "alpha"},
		"cudnnConvolutionBackwardData":            {10: "beta", 1: "alpha"},
		"cudnnSoftmaxForward":                     {6: "beta", 3: "alpha"},
		"cudnnSoftmaxBackward":                    {8: "beta", 3: "alpha"},
		"cudnnPoolingForward":                     {5: "beta", 2: "alpha"},
		"cudnnPoolingBackward":                    {9: "beta", 2: "alpha"},
		"cudnnActivationForward":                  {5: "beta", 2: "alpha"},
		"cudnnActivationBackward":                 {9: "beta", 2: "alpha"},
		"cudnnLRNCrossChannelForward":             {6: "beta", 3: "alpha"},
		"cudnnLRNCrossChannelBackward":            {10: "beta", 3: "alpha"},
		"cudnnDivisiveNormalizationForward":       {9: "beta", 3: "alpha"},
		"cudnnDivisiveNormalizationBackward":      {10: "beta", 3: "alpha"},
		"cudnnBatchNormalizationForwardTraining":  {3: "beta", 2: "alpha"},
		"cudnnBatchNormalizationForwardInference": {3: "beta", 2: "alpha"},
		"cudnnSpatialTfSamplerForward":            {6: "beta", 2: "alpha"},
		"cudnnSpatialTfSamplerBackward":           {5: "beta", 2: "alpha"},
	}

	creations = map[string]string{
		"cudnnHandle_t":                       "",
		"cudnnTensorDescriptor_t":             "",
		"cudnnOpTensorDescriptor_t":           "cudnnSetOpTensorDescriptor",
		"cudnnReduceTensorDescriptor_t":       "cudnnSetReduceTensorDescriptor",
		"cudnnFilterDescriptor_t":             "",
		"cudnnConvolutionDescriptor_t":        "",
		"cudnnPoolingDescriptor_t":            "",
		"cudnnActivationDescriptor_t":         "cudnnSetActivationDescriptor",
		"cudnnLRNDescriptor_t":                "cudnnSetLRNDescriptor",
		"cudnnSpatialTransformerDescriptor_t": "",
		"cudnnDropoutDescriptor_t":            "cudnnSetDropoutDescriptor",
		"cudnnRNNDescriptor_t":                "cudnnSetRNNDescriptor",
		"cudnnPersistentRNNPlan_t":            "cudnnSetPersistentRNNPlan",
		"cudnnCTCLossDescriptor_t":            "cudnnSetCTCLossDescriptor",
	}

	retVals = map[string]map[int]string{
		"cudnnGetProperty":                                   {1: "value"},
		"cudnnCreate":                                        {0: "handle"},
		"cudnnGetStream":                                     {1: "streamId"},
		"cudnnCreateTensorDescriptor":                        {0: "tensorDesc"},
		"cudnnGetTensor4dDescriptor":                         {9: "wStride"},
		"cudnnGetTensorNdDescriptor":                         {3: "nbDims"},
		"cudnnGetTensorSizeInBytes":                          {1: "size"},
		"cudnnCreateOpTensorDescriptor":                      {0: "opTensorDesc"},
		"cudnnGetOpTensorDescriptor":                         {3: "opTensorNanOpt"},
		"cudnnCreateReduceTensorDescriptor":                  {0: "reduceTensorDesc"},
		"cudnnGetReduceTensorDescriptor":                     {5: "reduceTensorIndicesType"},
		"cudnnGetReductionIndicesSize":                       {4: "sizeInBytes"},
		"cudnnGetReductionWorkspaceSize":                     {4: "sizeInBytes"},
		"cudnnCreateFilterDescriptor":                        {0: "filterDesc"},
		"cudnnGetFilter4dDescriptor":                         {6: "w"},
		"cudnnGetFilterNdDescriptor":                         {4: "nbDims"},
		"cudnnCreateConvolutionDescriptor":                   {0: "convDesc"},
		"cudnnGetConvolutionMathType":                        {1: "mathType"},
		"cudnnGetConvolutionGroupCount":                      {1: "groupCount"},
		"cudnnGetConvolution2dDescriptor":                    {8: "computeType"},
		"cudnnGetConvolution2dForwardOutputDim":              {6: "w"},
		"cudnnGetConvolutionNdDescriptor":                    {7: "computeType"},
		"cudnnGetConvolutionForwardAlgorithmMaxCount":        {1: "count"},
		"cudnnGetConvolutionForwardAlgorithm":                {7: "algo"},
		"cudnnGetConvolutionForwardAlgorithm_v7":             {7: "perfResults"},
		"cudnnGetConvolutionForwardWorkspaceSize":            {6: "sizeInBytes"},
		"cudnnGetConvolutionBackwardFilterAlgorithmMaxCount": {1: "count"},
		"cudnnGetConvolutionBackwardFilterAlgorithm":         {7: "algo"},
		"cudnnGetConvolutionBackwardFilterAlgorithm_v7":      {7: "perfResults"},
		"cudnnGetConvolutionBackwardFilterWorkspaceSize":     {6: "sizeInBytes"},
		"cudnnGetConvolutionBackwardDataAlgorithmMaxCount":   {1: "count"},
		"cudnnGetConvolutionBackwardDataAlgorithm":           {7: "algo"},
		"cudnnGetConvolutionBackwardDataAlgorithm_v7":        {7: "perfResults"},
		"cudnnGetConvolutionBackwardDataWorkspaceSize":       {6: "sizeInBytes"},
		"cudnnCreatePoolingDescriptor":                       {0: "poolingDesc"},
		"cudnnGetPooling2dDescriptor":                        {8: "horizontalStride"},
		"cudnnGetPoolingNdDescriptor":                        {4: "nbDims"},
		"cudnnGetPooling2dForwardOutputDim":                  {5: "w"},
		"cudnnCreateActivationDescriptor":                    {0: "activationDesc"},
		"cudnnGetActivationDescriptor":                       {3: "coef"},
		"cudnnCreateLRNDescriptor":                           {0: "normDesc"},
		"cudnnGetLRNDescriptor":                              {4: "lrnK"},
		"cudnnCreateSpatialTransformerDescriptor":            {0: "stDesc"},
		"cudnnCreateDropoutDescriptor":                       {0: "dropoutDesc"},
		"cudnnDropoutGetStatesSize":                          {1: "sizeInBytes"},
		"cudnnDropoutGetReserveSpaceSize":                    {1: "sizeInBytes"},
		"cudnnGetDropoutDescriptor":                          {4: "seed"},
		"cudnnCreateRNNDescriptor":                           {0: "rnnDesc"},
		"cudnnCreatePersistentRNNPlan":                       {3: "plan"},
		"cudnnGetRNNDescriptor":                              {9: "dataType"},
		"cudnnGetRNNWorkspaceSize":                           {4: "sizeInBytes"},
		"cudnnGetRNNTrainingReserveSize":                     {4: "sizeInBytes"},
		"cudnnGetRNNParamsSize":                              {3: "sizeInBytes"},
		"cudnnGetRNNLinLayerMatrixParams":                    {8: "linLayerMat"},
		"cudnnGetRNNLinLayerBiasParams":                      {8: "linLayerBias"},
		"cudnnCreateCTCLossDescriptor":                       {0: "ctcLossDesc"},
		"cudnnGetCTCLossDescriptor":                          {1: "compType"},
		"cudnnGetCTCLossWorkspaceSize":                       {8: "sizeInBytes"},
	}
}

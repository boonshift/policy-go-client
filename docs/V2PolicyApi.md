# \V2PolicyApi

All URIs are relative to *http://policy-service*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmendPolicy**](V2PolicyApi.md#AmendPolicy) | **Post** /v2/policy/amendpolicy/{id} | Amends an existing policy.
[**CancelPolicy**](V2PolicyApi.md#CancelPolicy) | **Post** /v2/policy/cancelpolicy/{id} | Cancel an existing lapsing policy.
[**ConvertQuotes**](V2PolicyApi.md#ConvertQuotes) | **Post** /v2/policy/convertquotes | Converts one or more quotes into policies.
[**CreateInvoice**](V2PolicyApi.md#CreateInvoice) | **Post** /v2/policy/invoice | Creates an individual invoice.
[**CreateInvoiceLine**](V2PolicyApi.md#CreateInvoiceLine) | **Post** /v2/policy/invoiceline | Creates an individual invoiceline.
[**CreatePaymentAttempt**](V2PolicyApi.md#CreatePaymentAttempt) | **Post** /v2/policy/paymentattempt | Creates an individual paymentattempt.
[**CreatePolicy**](V2PolicyApi.md#CreatePolicy) | **Post** /v2/policy/policy | Creates an individual policy.
[**CreatePolicyBundle**](V2PolicyApi.md#CreatePolicyBundle) | **Post** /v2/policy/policybundle | Creates an individual policybundle.
[**CreatePolicyHolder**](V2PolicyApi.md#CreatePolicyHolder) | **Post** /v2/policy/policyholder | Creates an individual policyholder.
[**CreatePolicyHolderNote**](V2PolicyApi.md#CreatePolicyHolderNote) | **Post** /v2/policy/policyholdernote | Creates an individual policyholdernote.
[**CreatePolicyNote**](V2PolicyApi.md#CreatePolicyNote) | **Post** /v2/policy/policynote | Creates an individual policynote.
[**CreateRenewalLink**](V2PolicyApi.md#CreateRenewalLink) | **Post** /v2/policy/renewal-link | Takes an existing proposal and creates a policy holder and renewal link for the first quote in the proposal.
[**DeleteInvoice**](V2PolicyApi.md#DeleteInvoice) | **Delete** /v2/policy/invoice/{id} | Deletes an individual invoice.
[**DeleteInvoiceLine**](V2PolicyApi.md#DeleteInvoiceLine) | **Delete** /v2/policy/invoiceline/{id} | Deletes an individual invoiceline.
[**DeletePaymentAttempt**](V2PolicyApi.md#DeletePaymentAttempt) | **Delete** /v2/policy/paymentattempt/{id} | Deletes an individual paymentattempt.
[**DeletePolicy**](V2PolicyApi.md#DeletePolicy) | **Delete** /v2/policy/policy/{id} | Deletes an individual policy.
[**DeletePolicyBundle**](V2PolicyApi.md#DeletePolicyBundle) | **Delete** /v2/policy/policybundle/{id} | Deletes an individual policybundle.
[**DeletePolicyHolder**](V2PolicyApi.md#DeletePolicyHolder) | **Delete** /v2/policy/policyholder/{id} | Deletes an individual policyholder.
[**DeletePolicyHolderNote**](V2PolicyApi.md#DeletePolicyHolderNote) | **Delete** /v2/policy/policyholdernote/{id} | Deletes an individual policyholdernote.
[**DeletePolicyNote**](V2PolicyApi.md#DeletePolicyNote) | **Delete** /v2/policy/policynote/{id} | Deletes an individual policynote.
[**EndorsePolicy**](V2PolicyApi.md#EndorsePolicy) | **Post** /v2/policy/endorsepolicy/{id} | Endorses an existing policy.
[**GetAllInvoiceLines**](V2PolicyApi.md#GetAllInvoiceLines) | **Get** /v2/policy/invoiceline | Gets all invoicelines visible to the logged in user.
[**GetAllInvoices**](V2PolicyApi.md#GetAllInvoices) | **Get** /v2/policy/invoice | Gets all invoices visible to the logged in user.
[**GetAllPaymentAttempts**](V2PolicyApi.md#GetAllPaymentAttempts) | **Get** /v2/policy/paymentattempt | Gets all paymentattempts visible to the logged in user.
[**GetAllPolicies**](V2PolicyApi.md#GetAllPolicies) | **Get** /v2/policy/policy | Gets all policies visible to the logged in user.
[**GetAllPolicyBundles**](V2PolicyApi.md#GetAllPolicyBundles) | **Get** /v2/policy/policybundle | Gets all policybundles visible to the logged in user.
[**GetAllPolicyHolderNotes**](V2PolicyApi.md#GetAllPolicyHolderNotes) | **Get** /v2/policy/policyholdernote | Gets all policyholdernotes visible to the logged in user.
[**GetAllPolicyHolders**](V2PolicyApi.md#GetAllPolicyHolders) | **Get** /v2/policy/policyholder | Gets all policyholders visible to the logged in user.
[**GetAllPolicyNotes**](V2PolicyApi.md#GetAllPolicyNotes) | **Get** /v2/policy/policynote | Gets all policynotes visible to the logged in user.
[**GetInvoice**](V2PolicyApi.md#GetInvoice) | **Get** /v2/policy/invoice/{id} | Gets an individual invoice.
[**GetInvoiceLine**](V2PolicyApi.md#GetInvoiceLine) | **Get** /v2/policy/invoiceline/{id} | Gets an individual invoiceline.
[**GetInvoiceLinesByInvoiceId**](V2PolicyApi.md#GetInvoiceLinesByInvoiceId) | **Get** /v2/policy/invoiceline/invoiceid/{invoiceid} | Gets all invoicelines visible to the logged in user with the specified invoiceid.
[**GetInvoiceLinesByPolicyHolderId**](V2PolicyApi.md#GetInvoiceLinesByPolicyHolderId) | **Get** /v2/policy/invoiceline/policyholderid/{policyholderid} | Gets all invoicelines visible to the logged in user with the specified policyholderid.
[**GetInvoiceLinesByPolicyId**](V2PolicyApi.md#GetInvoiceLinesByPolicyId) | **Get** /v2/policy/invoiceline/policyid/{policyid} | Gets all invoicelines visible to the logged in user with the specified policyid.
[**GetInvoicesByPolicyHolderId**](V2PolicyApi.md#GetInvoicesByPolicyHolderId) | **Get** /v2/policy/invoice/policyholderid/{policyholderid} | Gets all invoices visible to the logged in user with the specified policyholderid.
[**GetPaymentAttempt**](V2PolicyApi.md#GetPaymentAttempt) | **Get** /v2/policy/paymentattempt/{id} | Gets an individual paymentattempt.
[**GetPaymentAttemptsByInvoiceId**](V2PolicyApi.md#GetPaymentAttemptsByInvoiceId) | **Get** /v2/policy/paymentattempt/invoiceid/{invoiceid} | Gets all paymentattempts visible to the logged in user with the specified invoiceid.
[**GetPoliciesByPolicyBundleId**](V2PolicyApi.md#GetPoliciesByPolicyBundleId) | **Get** /v2/policy/policy/policybundleid/{policybundleid} | Gets all policies visible to the logged in user with the specified policybundleid.
[**GetPoliciesByPolicyHolderId**](V2PolicyApi.md#GetPoliciesByPolicyHolderId) | **Get** /v2/policy/policy/policyholderid/{policyholderid} | Gets all policies visible to the logged in user with the specified policyholderid.
[**GetPoliciesByPolicyNumber**](V2PolicyApi.md#GetPoliciesByPolicyNumber) | **Get** /v2/policy/policy/policynumber/{policynumber} | Gets all policies visible to the logged in user with the specified policynumber.
[**GetPolicy**](V2PolicyApi.md#GetPolicy) | **Get** /v2/policy/policy/{id} | Gets an individual policy.
[**GetPolicyBundle**](V2PolicyApi.md#GetPolicyBundle) | **Get** /v2/policy/policybundle/{id} | Gets an individual policybundle.
[**GetPolicyHistory**](V2PolicyApi.md#GetPolicyHistory) | **Get** /v2/policy/policy/{id}/history | Get the history of the policy.
[**GetPolicyHolder**](V2PolicyApi.md#GetPolicyHolder) | **Get** /v2/policy/policyholder/{id} | Gets an individual policyholder.
[**GetPolicyHolderNote**](V2PolicyApi.md#GetPolicyHolderNote) | **Get** /v2/policy/policyholdernote/{id} | Gets an individual policyholdernote.
[**GetPolicyHolderNotesByPolicyHolderId**](V2PolicyApi.md#GetPolicyHolderNotesByPolicyHolderId) | **Get** /v2/policy/policyholdernote/policyholderid/{policyholderid} | Gets all policyholdernotes visible to the logged in user with the specified policyholderid.
[**GetPolicyHoldersByUserId**](V2PolicyApi.md#GetPolicyHoldersByUserId) | **Get** /v2/policy/policyholder/userid/{userid} | Gets all policyholders visible to the logged in user with the specified userid.
[**GetPolicyNote**](V2PolicyApi.md#GetPolicyNote) | **Get** /v2/policy/policynote/{id} | Gets an individual policynote.
[**GetPolicyNotesByPolicyId**](V2PolicyApi.md#GetPolicyNotesByPolicyId) | **Get** /v2/policy/policynote/policyid/{policyid} | Gets all policynotes visible to the logged in user with the specified policyid.
[**LapsePolicy**](V2PolicyApi.md#LapsePolicy) | **Post** /v2/policy/lapsepolicy/{id} | Lapse an existing lapsing policy.
[**ManuallyPayInvoice**](V2PolicyApi.md#ManuallyPayInvoice) | **Post** /v2/policy/invoice/manualpay/{id} | Initiate the manual payment of an invoice.
[**PayInvoice**](V2PolicyApi.md#PayInvoice) | **Post** /v2/policy/invoice/pay/{id} | Initiate payment of an invoice.
[**ReviewPolicy**](V2PolicyApi.md#ReviewPolicy) | **Post** /v2/policy/reviewpolicy/{id} | Approves and/or revises policy stuck in review.
[**SetInvoiceLineOwner**](V2PolicyApi.md#SetInvoiceLineOwner) | **Put** /v2/policy/invoiceline/{id}/owner | Changes the owner of the invoiceline.
[**SetInvoiceOwner**](V2PolicyApi.md#SetInvoiceOwner) | **Put** /v2/policy/invoice/{id}/owner | Changes the owner of the invoice.
[**SetPolicyBundleOwner**](V2PolicyApi.md#SetPolicyBundleOwner) | **Put** /v2/policy/policybundle/{id}/owner | Changes the owner of the policybundle.
[**SetPolicyHolderOwner**](V2PolicyApi.md#SetPolicyHolderOwner) | **Put** /v2/policy/policyholder/{id}/owner | Changes the owner of the policyholder.
[**SetPolicyOwner**](V2PolicyApi.md#SetPolicyOwner) | **Put** /v2/policy/policy/{id}/owner | Changes the owner of the policy.
[**TerminatePolicy**](V2PolicyApi.md#TerminatePolicy) | **Post** /v2/policy/terminatepolicy/{id} | Terminate an existing lapsing policy.
[**UpdateInvoice**](V2PolicyApi.md#UpdateInvoice) | **Put** /v2/policy/invoice/{id} | Updates an individual invoice.
[**UpdateInvoiceLine**](V2PolicyApi.md#UpdateInvoiceLine) | **Put** /v2/policy/invoiceline/{id} | Updates an individual invoiceline.
[**UpdatePaymentAttempt**](V2PolicyApi.md#UpdatePaymentAttempt) | **Put** /v2/policy/paymentattempt/{id} | Updates an individual paymentattempt.
[**UpdatePolicy**](V2PolicyApi.md#UpdatePolicy) | **Put** /v2/policy/policy/{id} | Updates an individual policy.
[**UpdatePolicyBundle**](V2PolicyApi.md#UpdatePolicyBundle) | **Put** /v2/policy/policybundle/{id} | Updates an individual policybundle.
[**UpdatePolicyHolder**](V2PolicyApi.md#UpdatePolicyHolder) | **Put** /v2/policy/policyholder/{id} | Updates an individual policyholder.
[**UpdatePolicyHolderNote**](V2PolicyApi.md#UpdatePolicyHolderNote) | **Put** /v2/policy/policyholdernote/{id} | Updates an individual policyholdernote.
[**UpdatePolicyNote**](V2PolicyApi.md#UpdatePolicyNote) | **Put** /v2/policy/policynote/{id} | Updates an individual policynote.



## AmendPolicy

> SuccessResultAmendPolicyResponseModel AmendPolicy(ctx, id, amendPolicyModel)

Amends an existing policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**amendPolicyModel** | [**AmendPolicyModel**](AmendPolicyModel.md)|  | 

### Return type

[**SuccessResultAmendPolicyResponseModel**](SuccessResultAmendPolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPolicy

> SuccessResultCancelPolicyResponseModel CancelPolicy(ctx, id, policyCancelModel)

Cancel an existing lapsing policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyCancelModel** | [**PolicyCancelModel**](PolicyCancelModel.md)|  | 

### Return type

[**SuccessResultCancelPolicyResponseModel**](SuccessResultCancelPolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertQuotes

> SuccessResultConvertQuoteResponse ConvertQuotes(ctx, convertQuoteModel)

Converts one or more quotes into policies.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**convertQuoteModel** | [**ConvertQuoteModel**](ConvertQuoteModel.md)|  | 

### Return type

[**SuccessResultConvertQuoteResponse**](SuccessResultConvertQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoice

> SuccessResultInvoice CreateInvoice(ctx, invoiceCreateModel)

Creates an individual invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceCreateModel** | [**InvoiceCreateModel**](InvoiceCreateModel.md)|  | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceLine

> SuccessResultInvoiceLine CreateInvoiceLine(ctx, invoiceLineCreateModel)

Creates an individual invoiceline.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceLineCreateModel** | [**InvoiceLineCreateModel**](InvoiceLineCreateModel.md)|  | 

### Return type

[**SuccessResultInvoiceLine**](SuccessResultInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentAttempt

> SuccessResultPaymentAttempt CreatePaymentAttempt(ctx, paymentAttemptCreateModel)

Creates an individual paymentattempt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentAttemptCreateModel** | [**PaymentAttemptCreateModel**](PaymentAttemptCreateModel.md)|  | 

### Return type

[**SuccessResultPaymentAttempt**](SuccessResultPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> SuccessResultPolicy CreatePolicy(ctx, policyCreateModel)

Creates an individual policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyCreateModel** | [**PolicyCreateModel**](PolicyCreateModel.md)|  | 

### Return type

[**SuccessResultPolicy**](SuccessResultPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyBundle

> SuccessResultPolicyBundle CreatePolicyBundle(ctx, policyBundleCreateModel)

Creates an individual policybundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyBundleCreateModel** | [**PolicyBundleCreateModel**](PolicyBundleCreateModel.md)|  | 

### Return type

[**SuccessResultPolicyBundle**](SuccessResultPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyHolder

> SuccessResultPolicyHolder CreatePolicyHolder(ctx, policyHolderCreateModel)

Creates an individual policyholder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyHolderCreateModel** | [**PolicyHolderCreateModel**](PolicyHolderCreateModel.md)|  | 

### Return type

[**SuccessResultPolicyHolder**](SuccessResultPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyHolderNote

> SuccessResultPolicyHolderNote CreatePolicyHolderNote(ctx, policyHolderNoteCreateModel)

Creates an individual policyholdernote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyHolderNoteCreateModel** | [**PolicyHolderNoteCreateModel**](PolicyHolderNoteCreateModel.md)|  | 

### Return type

[**SuccessResultPolicyHolderNote**](SuccessResultPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyNote

> SuccessResultPolicyNote CreatePolicyNote(ctx, policyNoteCreateModel)

Creates an individual policynote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyNoteCreateModel** | [**PolicyNoteCreateModel**](PolicyNoteCreateModel.md)|  | 

### Return type

[**SuccessResultPolicyNote**](SuccessResultPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRenewalLink

> SuccessResultRenewalLinkResponseModel CreateRenewalLink(ctx, renewalLinkModel)

Takes an existing proposal and creates a policy holder and renewal link for the first quote in the proposal.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**renewalLinkModel** | [**RenewalLinkModel**](RenewalLinkModel.md)|  | 

### Return type

[**SuccessResultRenewalLinkResponseModel**](SuccessResultRenewalLinkResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvoice

> SuccessResultInvoice DeleteInvoice(ctx, id)

Deletes an individual invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvoiceLine

> SuccessResultInvoiceLine DeleteInvoiceLine(ctx, id)

Deletes an individual invoiceline.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoiceline ID. | 

### Return type

[**SuccessResultInvoiceLine**](SuccessResultInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentAttempt

> SuccessResultPaymentAttempt DeletePaymentAttempt(ctx, id)

Deletes an individual paymentattempt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The paymentattempt ID. | 

### Return type

[**SuccessResultPaymentAttempt**](SuccessResultPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> SuccessResultPolicy DeletePolicy(ctx, id)

Deletes an individual policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 

### Return type

[**SuccessResultPolicy**](SuccessResultPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyBundle

> SuccessResultPolicyBundle DeletePolicyBundle(ctx, id)

Deletes an individual policybundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policybundle ID. | 

### Return type

[**SuccessResultPolicyBundle**](SuccessResultPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyHolder

> SuccessResultPolicyHolder DeletePolicyHolder(ctx, id)

Deletes an individual policyholder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholder ID. | 

### Return type

[**SuccessResultPolicyHolder**](SuccessResultPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyHolderNote

> SuccessResultPolicyHolderNote DeletePolicyHolderNote(ctx, id)

Deletes an individual policyholdernote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholdernote ID. | 

### Return type

[**SuccessResultPolicyHolderNote**](SuccessResultPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyNote

> SuccessResultPolicyNote DeletePolicyNote(ctx, id)

Deletes an individual policynote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policynote ID. | 

### Return type

[**SuccessResultPolicyNote**](SuccessResultPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndorsePolicy

> SuccessResultEndorsePolicyResponseModel EndorsePolicy(ctx, id, policyEndorseModel)

Endorses an existing policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyEndorseModel** | [**PolicyEndorseModel**](PolicyEndorseModel.md)|  | 

### Return type

[**SuccessResultEndorsePolicyResponseModel**](SuccessResultEndorsePolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInvoiceLines

> SuccessResultListInvoiceLine GetAllInvoiceLines(ctx, optional)

Gets all invoicelines visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInvoiceLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInvoiceLinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoiceLine**](SuccessResultListInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInvoices

> SuccessResultListInvoice GetAllInvoices(ctx, optional)

Gets all invoices visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoice**](SuccessResultListInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPaymentAttempts

> SuccessResultListPaymentAttempt GetAllPaymentAttempts(ctx, optional)

Gets all paymentattempts visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPaymentAttemptsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPaymentAttemptsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPaymentAttempt**](SuccessResultListPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPolicies

> SuccessResultListPolicy GetAllPolicies(ctx, optional)

Gets all policies visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicy**](SuccessResultListPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPolicyBundles

> SuccessResultListPolicyBundle GetAllPolicyBundles(ctx, optional)

Gets all policybundles visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPolicyBundlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPolicyBundlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyBundle**](SuccessResultListPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPolicyHolderNotes

> SuccessResultListPolicyHolderNote GetAllPolicyHolderNotes(ctx, optional)

Gets all policyholdernotes visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPolicyHolderNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPolicyHolderNotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyHolderNote**](SuccessResultListPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPolicyHolders

> SuccessResultListPolicyHolder GetAllPolicyHolders(ctx, optional)

Gets all policyholders visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPolicyHoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPolicyHoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyHolder**](SuccessResultListPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPolicyNotes

> SuccessResultListPolicyNote GetAllPolicyNotes(ctx, optional)

Gets all policynotes visible to the logged in user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllPolicyNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllPolicyNotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyNote**](SuccessResultListPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> SuccessResultInvoice GetInvoice(ctx, id)

Gets an individual invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLine

> SuccessResultInvoiceLine GetInvoiceLine(ctx, id)

Gets an individual invoiceline.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoiceline ID. | 

### Return type

[**SuccessResultInvoiceLine**](SuccessResultInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLinesByInvoiceId

> SuccessResultListInvoiceLine GetInvoiceLinesByInvoiceId(ctx, invoiceid, optional)

Gets all invoicelines visible to the logged in user with the specified invoiceid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceid** | **string**| The invoiceId of the invoiceline. | 
 **optional** | ***GetInvoiceLinesByInvoiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoiceLinesByInvoiceIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoiceLine**](SuccessResultListInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLinesByPolicyHolderId

> SuccessResultListInvoiceLine GetInvoiceLinesByPolicyHolderId(ctx, policyholderid, optional)

Gets all invoicelines visible to the logged in user with the specified policyholderid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyholderid** | **string**| The policyHolderId of the invoiceline. | 
 **optional** | ***GetInvoiceLinesByPolicyHolderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoiceLinesByPolicyHolderIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoiceLine**](SuccessResultListInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLinesByPolicyId

> SuccessResultListInvoiceLine GetInvoiceLinesByPolicyId(ctx, policyid, optional)

Gets all invoicelines visible to the logged in user with the specified policyid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyid** | **string**| The policyId of the invoiceline. | 
 **optional** | ***GetInvoiceLinesByPolicyIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoiceLinesByPolicyIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoiceLine**](SuccessResultListInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicesByPolicyHolderId

> SuccessResultListInvoice GetInvoicesByPolicyHolderId(ctx, policyholderid, optional)

Gets all invoices visible to the logged in user with the specified policyholderid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyholderid** | **string**| The policyHolderId of the invoice. | 
 **optional** | ***GetInvoicesByPolicyHolderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoicesByPolicyHolderIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListInvoice**](SuccessResultListInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentAttempt

> SuccessResultPaymentAttempt GetPaymentAttempt(ctx, id)

Gets an individual paymentattempt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The paymentattempt ID. | 

### Return type

[**SuccessResultPaymentAttempt**](SuccessResultPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentAttemptsByInvoiceId

> SuccessResultListPaymentAttempt GetPaymentAttemptsByInvoiceId(ctx, invoiceid, optional)

Gets all paymentattempts visible to the logged in user with the specified invoiceid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceid** | **string**| The invoiceId of the paymentattempt. | 
 **optional** | ***GetPaymentAttemptsByInvoiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentAttemptsByInvoiceIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPaymentAttempt**](SuccessResultListPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoliciesByPolicyBundleId

> SuccessResultListPolicy GetPoliciesByPolicyBundleId(ctx, policybundleid, optional)

Gets all policies visible to the logged in user with the specified policybundleid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policybundleid** | **string**| The policyBundleId of the policy. | 
 **optional** | ***GetPoliciesByPolicyBundleIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPoliciesByPolicyBundleIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicy**](SuccessResultListPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoliciesByPolicyHolderId

> SuccessResultListPolicy GetPoliciesByPolicyHolderId(ctx, policyholderid, optional)

Gets all policies visible to the logged in user with the specified policyholderid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyholderid** | **string**| The policyHolderId of the policy. | 
 **optional** | ***GetPoliciesByPolicyHolderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPoliciesByPolicyHolderIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicy**](SuccessResultListPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoliciesByPolicyNumber

> SuccessResultListPolicy GetPoliciesByPolicyNumber(ctx, policynumber, optional)

Gets all policies visible to the logged in user with the specified policynumber.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policynumber** | **string**| The policyNumber of the policy. | 
 **optional** | ***GetPoliciesByPolicyNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPoliciesByPolicyNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicy**](SuccessResultListPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> SuccessResultPolicy GetPolicy(ctx, id)

Gets an individual policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 

### Return type

[**SuccessResultPolicy**](SuccessResultPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyBundle

> SuccessResultPolicyBundle GetPolicyBundle(ctx, id)

Gets an individual policybundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policybundle ID. | 

### Return type

[**SuccessResultPolicyBundle**](SuccessResultPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyHistory

> SuccessResultListPolicyHistory GetPolicyHistory(ctx, id, optional)

Get the history of the policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
 **optional** | ***GetPolicyHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyHistory**](SuccessResultListPolicyHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyHolder

> SuccessResultPolicyHolder GetPolicyHolder(ctx, id)

Gets an individual policyholder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholder ID. | 

### Return type

[**SuccessResultPolicyHolder**](SuccessResultPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyHolderNote

> SuccessResultPolicyHolderNote GetPolicyHolderNote(ctx, id)

Gets an individual policyholdernote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholdernote ID. | 

### Return type

[**SuccessResultPolicyHolderNote**](SuccessResultPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyHolderNotesByPolicyHolderId

> SuccessResultListPolicyHolderNote GetPolicyHolderNotesByPolicyHolderId(ctx, policyholderid, optional)

Gets all policyholdernotes visible to the logged in user with the specified policyholderid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyholderid** | **string**| The policyHolderId of the policyholdernote. | 
 **optional** | ***GetPolicyHolderNotesByPolicyHolderIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyHolderNotesByPolicyHolderIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyHolderNote**](SuccessResultListPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyHoldersByUserId

> SuccessResultListPolicyHolder GetPolicyHoldersByUserId(ctx, userid, optional)

Gets all policyholders visible to the logged in user with the specified userid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string**| The userId of the policyholder. | 
 **optional** | ***GetPolicyHoldersByUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyHoldersByUserIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyHolder**](SuccessResultListPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyNote

> SuccessResultPolicyNote GetPolicyNote(ctx, id)

Gets an individual policynote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policynote ID. | 

### Return type

[**SuccessResultPolicyNote**](SuccessResultPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyNotesByPolicyId

> SuccessResultListPolicyNote GetPolicyNotesByPolicyId(ctx, policyid, optional)

Gets all policynotes visible to the logged in user with the specified policyid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyid** | **string**| The policyId of the policynote. | 
 **optional** | ***GetPolicyNotesByPolicyIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyNotesByPolicyIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| requested offset in requested collection | 
 **limit** | **optional.Int32**| limit of results in requested collection | 
 **sort** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**SuccessResultListPolicyNote**](SuccessResultListPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LapsePolicy

> SuccessResultLapsePolicyResponseModel LapsePolicy(ctx, id, policyLapseModel)

Lapse an existing lapsing policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyLapseModel** | [**PolicyLapseModel**](PolicyLapseModel.md)|  | 

### Return type

[**SuccessResultLapsePolicyResponseModel**](SuccessResultLapsePolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManuallyPayInvoice

> SuccessResultInvoice ManuallyPayInvoice(ctx, id, body)

Initiate the manual payment of an invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 
**body** | **map[string]interface{}**|  | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayInvoice

> SuccessResultInvoice PayInvoice(ctx, id, body)

Initiate payment of an invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 
**body** | **map[string]interface{}**|  | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewPolicy

> SuccessResultReviewPolicyResponseModel ReviewPolicy(ctx, id, policyReviewModel)

Approves and/or revises policy stuck in review.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyReviewModel** | [**PolicyReviewModel**](PolicyReviewModel.md)|  | 

### Return type

[**SuccessResultReviewPolicyResponseModel**](SuccessResultReviewPolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInvoiceLineOwner

> SuccessResultInvoiceLine SetInvoiceLineOwner(ctx, id, setInvoiceLineOwnerModel)

Changes the owner of the invoiceline.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoiceline ID. | 
**setInvoiceLineOwnerModel** | [**SetInvoiceLineOwnerModel**](SetInvoiceLineOwnerModel.md)|  | 

### Return type

[**SuccessResultInvoiceLine**](SuccessResultInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInvoiceOwner

> SuccessResultInvoice SetInvoiceOwner(ctx, id, setInvoiceOwnerModel)

Changes the owner of the invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 
**setInvoiceOwnerModel** | [**SetInvoiceOwnerModel**](SetInvoiceOwnerModel.md)|  | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicyBundleOwner

> SuccessResultPolicyBundle SetPolicyBundleOwner(ctx, id, setPolicyBundleOwnerModel)

Changes the owner of the policybundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policybundle ID. | 
**setPolicyBundleOwnerModel** | [**SetPolicyBundleOwnerModel**](SetPolicyBundleOwnerModel.md)|  | 

### Return type

[**SuccessResultPolicyBundle**](SuccessResultPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicyHolderOwner

> SuccessResultPolicyHolder SetPolicyHolderOwner(ctx, id, setPolicyHolderOwnerModel)

Changes the owner of the policyholder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholder ID. | 
**setPolicyHolderOwnerModel** | [**SetPolicyHolderOwnerModel**](SetPolicyHolderOwnerModel.md)|  | 

### Return type

[**SuccessResultPolicyHolder**](SuccessResultPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPolicyOwner

> SuccessResultPolicy SetPolicyOwner(ctx, id, setPolicyOwnerModel)

Changes the owner of the policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**setPolicyOwnerModel** | [**SetPolicyOwnerModel**](SetPolicyOwnerModel.md)|  | 

### Return type

[**SuccessResultPolicy**](SuccessResultPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminatePolicy

> SuccessResultTerminatePolicyResponseModel TerminatePolicy(ctx, id, policyTerminateModel)

Terminate an existing lapsing policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyTerminateModel** | [**PolicyTerminateModel**](PolicyTerminateModel.md)|  | 

### Return type

[**SuccessResultTerminatePolicyResponseModel**](SuccessResultTerminatePolicyResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> SuccessResultInvoice UpdateInvoice(ctx, id, invoiceUpdateModel)

Updates an individual invoice.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoice ID. | 
**invoiceUpdateModel** | [**InvoiceUpdateModel**](InvoiceUpdateModel.md)|  | 

### Return type

[**SuccessResultInvoice**](SuccessResultInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceLine

> SuccessResultInvoiceLine UpdateInvoiceLine(ctx, id, invoiceLineUpdateModel)

Updates an individual invoiceline.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The invoiceline ID. | 
**invoiceLineUpdateModel** | [**InvoiceLineUpdateModel**](InvoiceLineUpdateModel.md)|  | 

### Return type

[**SuccessResultInvoiceLine**](SuccessResultInvoiceLine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentAttempt

> SuccessResultPaymentAttempt UpdatePaymentAttempt(ctx, id, paymentAttemptUpdateModel)

Updates an individual paymentattempt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The paymentattempt ID. | 
**paymentAttemptUpdateModel** | [**PaymentAttemptUpdateModel**](PaymentAttemptUpdateModel.md)|  | 

### Return type

[**SuccessResultPaymentAttempt**](SuccessResultPaymentAttempt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> SuccessResultPolicy UpdatePolicy(ctx, id, policyUpdateModel)

Updates an individual policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policy ID. | 
**policyUpdateModel** | [**PolicyUpdateModel**](PolicyUpdateModel.md)|  | 

### Return type

[**SuccessResultPolicy**](SuccessResultPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyBundle

> SuccessResultPolicyBundle UpdatePolicyBundle(ctx, id, policyBundleUpdateModel)

Updates an individual policybundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policybundle ID. | 
**policyBundleUpdateModel** | [**PolicyBundleUpdateModel**](PolicyBundleUpdateModel.md)|  | 

### Return type

[**SuccessResultPolicyBundle**](SuccessResultPolicyBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyHolder

> SuccessResultPolicyHolder UpdatePolicyHolder(ctx, id, policyHolderUpdateModel)

Updates an individual policyholder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholder ID. | 
**policyHolderUpdateModel** | [**PolicyHolderUpdateModel**](PolicyHolderUpdateModel.md)|  | 

### Return type

[**SuccessResultPolicyHolder**](SuccessResultPolicyHolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyHolderNote

> SuccessResultPolicyHolderNote UpdatePolicyHolderNote(ctx, id, policyHolderNoteUpdateModel)

Updates an individual policyholdernote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policyholdernote ID. | 
**policyHolderNoteUpdateModel** | [**PolicyHolderNoteUpdateModel**](PolicyHolderNoteUpdateModel.md)|  | 

### Return type

[**SuccessResultPolicyHolderNote**](SuccessResultPolicyHolderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyNote

> SuccessResultPolicyNote UpdatePolicyNote(ctx, id, policyNoteUpdateModel)

Updates an individual policynote.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The policynote ID. | 
**policyNoteUpdateModel** | [**PolicyNoteUpdateModel**](PolicyNoteUpdateModel.md)|  | 

### Return type

[**SuccessResultPolicyNote**](SuccessResultPolicyNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


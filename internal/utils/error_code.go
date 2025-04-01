package utils

func ParseErrorCode(v string) string {
	switch v {
	case "1001":
		return "Default error code"
	case "1002":
		return "No MSISDN specified"
	case "1003":
		return "Transaction number is illegal"
	case "1004":
		return "The password is incorrect"
	case "1005":
		return "The provider does not exist"
	case "1006":
		return "The price plan with code {0} does not exist"
	case "1007":
		return "The subscriber does not have this individual price plan"
	case "1008":
		return "The subscriber has already ordered the individual price plan"
	case "1009":
		return "The subscription plan and provider are not a valid combination"
	case "1010":
		return "Language does not exist in the system"
	case "1011":
		return "The subscription plan doesn't exist"
	case "1012":
		return "The related product with code{0} does not exist"
	case "1013":
		return "The related product with code {0} does not belong to the subscription plan"
	case "1014":
		return "The price plan does not match current main product."
	case "1015":
		return "InitialState <initial state> is not valid"
	case "1016":
		return "The price plan is not a subscription default price plan"
	case "1017":
		return "The subscription default price plan does not exist"
	case "1018":
		return "The subscriber has not ordered this(these) product(s), deletion is not allowed"
	case "1019":
		return "The subscriber has not ordered this(these) product(s), activation is not allowed"
	case "1020":
		return "The attribute with code {0} does not exist"
	case "1021":
		return "The product specification does not have the attribute with code {0}"
	case "1022":
		return "The new subscription plan is same with the current one"
	case "1023":
		return "The subscription plan is not in effective period"
	case "1024":
		return "The new ICCID is the same as the old one,don't duplicate SIM card"
	case "1025":
		return "The new IMSI is the same as the old one, don't duplicate SIM card"
	case "1026":
		return "The product does not exist"
	case "1027":
		return "The old password is incorrect"
	case "1028":
		return "New MSISDN already exists in the system"
	case "1029":
		return "Account code does not exist"
	case "1030":
		return "The subscriber does't have the FnF service"
	case "1031":
		return "The format of fellow number is illegal"
	case "1032":
		return "The fellow number type doesn't exist"
	case "1033":
		return "Balance is not enough"
	case "1034":
		return "The fellow number is already in the list"
	case "1035":
		return "Exceeds the maximum operation"
	case "1036":
		return "The fellow number doesn't exist in the list"
	case "1037":
		return "The subscriber has not been two-way blocked"
	case "1038":
		return "The subscriber is suspended under request"
	case "1039":
		return "The subscriber has not been suspended under request"
	case "1040":
		return "The subscriber is blocked for peccancy"
	case "1041":
		return "The subscriber has been reported loss"
	case "1042":
		return "The subscriber not reported loss can not be restored"
	case "1043":
		return "The subscriber has not been blocked for peccancy"
	case "1044":
		return "The subscriber is not suspended under request"
	case "1045":
		return "The subscriber is two-way blocked and can not be one-way reactivated"
	case "1046":
		return "The subscriber is not blocked for debt"
	case "1047":
		return "The subscriber not in the blacklist can not release the blacklist"
	case "1048":
		return "The subscriber is two-way blocked for peccancy and can not be one-way reactivated for peccancy"
	case "1049":
		return "Balance type doesn`t exist in the system"
	case "1050":
		return "The recharge amount exceed the limit of this account balance type"
	case "1051":
		return "The recharge amount exceed the ceiling limit of this independent product"
	case "1052":
		return "The account status prohibits recharge for MSISDN "
	case "1053":
		return "Extension days can not be longger than 999"
	case "1054":
		return "EffDate can not be set to a date in the past"
	case "1055":
		return "ExpDate can not be set to a date in the past"
	case "1056":
		return "Unable to calculate expiry periods"
	case "1057":
		return "Subscriber is not a prepaid number"
	case "1058":
		return "Action is not valid"
	case "1059":
		return "EffType is not valid"
	case "1060":
		return "Accumulation  does not exist"
	case "1061":
		return "Voucher Center is busy"
	case "1062":
		return "The request data can not be parsed by Voucher Center"
	case "1063":
		return "The password of voucher card is invalid"
	case "1064":
		return "The voucher card is invalid"
	case "1065":
		return "The voucher card has been used"
	case "1066":
		return "The voucher card has been locked"
	case "1201":
		return "Loan amount is not in the list"
	case "1202":
		return "Loan amount exceeds the limitation"
	case "1203":
		return "Exceeds the maximum loan times of one period"
	case "1204":
		return "The balance is not in the allowed range"
	case "1205":
		return "Last loan records are not paid off"
	case "1206":
		return "The loan is not permitted for subssriber`s brand"
	case "1207":
		return "Subscriber is not in active state"
	case "1208":
		return "The activation period is less than the limitation"
	case "1209":
		return "The loan is not permitted for subssriber`s price plan"
	case "1210":
		return "On-net consumption in a period is not enough for loan"
	case "1211":
		return "Recharge times are not enough for loan"
	case "1212":
		return "The last loan is not return in 30 days"
	case "1213":
		return "The recharge amount of last month is not enough to loan"
	case "1301":
		return "The account of transfer-out is invalid"
	case "1302":
		return "The balance of transfer-out is insufficient"
	case "1303":
		return "The balance of transfer-in exceeds the maximum limitation"
	case "1304":
		return "Fail to transfer, because the transfer-out subscriber is forbidden to transfer"
	case "1305":
		return "The transfer amount is less than the minimum limitation."
	case "1306":
		return "Fail to transfer, because the price plan of the transfer-out subscriber is not allowed to transfer"
	case "1307":
		return "The transfer amount is larger than maximum limitation"
	case "1308":
		return "The account of transfer-in is invalid"
	case "1309":
		return "Fail to transfer, because the price plan of the transfer-out subscriber is not allowed to transfer"
	case "1310":
		return "Fail to transfer, because the product state of transfer-in subscriber is not allowed to transfer"
	case "1311":
		return "Fail to transfer, because the balance after transfering is less than the specified minimum balance"
	case "1312":
		return "Fail to transfer, because the transfer times in  a period(daily/weekly/monthly) is greater than the maximum transfer times"
	case "1313":
		return "The onnet age of transfer-out is not long enough to transfer"
	case "1314":
		return "Fail to transfer, because the daily transfer-out amount is greater than the ceilling limit"
	case "1315":
		return "Fail to transfer, because the daily transfer-in amount is greater than the ceilling limit"
	case "1316":
		return "The state of transfer-in subscriber is blocked"
	case "1317":
		return "The amount of transfer-out subscriber exceeds the monthly limitation"
	case "1318":
		return "The balance type of transfer-out subscriber doesn`t exist"
	case "1319":
		return "The balance type is not allowed to transfer"
	case "1320":
		return "The transfer-in subscriber is postpaid"
	case "1321":
		return "Transfer out subscriber exceeds the credit limit"
	case "1322":
		return "The msisdn of transfer in and transfer out can not be same"
	case "1323":
		return "The amount of transfer in subscriber exceeds the maximum limitation"
	case "1324":
		return "The transfer out subscriber is in the transfer blacklist, not allowed to transfer out"
	case "1325":
		return "The transfer-in amount is mandatory"
	case "1326":
		return "The transfer-in amount can not be positive or zero"
	case "1327":
		return "Fail to transfer, because the price plan of the transfer-in subscriber is not allowed to transfer"
	case "1328":
		return "The count of numbers exceeds the maximum count allowed by the fellow number nominations"
	case "1329":
		return "The subscriber has already ordered this(these) product(s), duplicate order is not allowed"
	case "9999":
		return "Pre-defined text"
	default:
		return "No description"
	}
}

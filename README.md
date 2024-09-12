# Don't forget to load your .env file

export $(<.env)

_Customer Interaction Management_

The basic features are listed below, so we will need these APIs to interface if available.

- Check customer status to see if the customer is subscribed to the service or not
- Subscription type
- Subscribe a customer
- Unsubscribe a customer
- Billing trace.

PREDICTION CHOOSE A TEAM
`Credit Goal: {home}-{away}! Gagnez {credit}F credit a chaque but de votre equipe si elle gagne le match! Envoyes {home-code} ou {away-code} par SMS au {sdc}. {price}{currency}/sms`

MT (WELCOME)
`Votre participation a ete enregistree. Si {team} gagne et marque des buts lors du prochain match, vous recevrez {price}{currency} deb bonus par but. {price}{currency}/souscription`

USER WIN
`Credit Goal: Felicitations! Le score final du match {home}-{away} est {score}. Votre compte va etre credite dans un delai de 72H de {price}{currency}`

USER LOSES
`?`

MT UNSUB
``

## Pseudo Code

1. User Press #101#36# on USSD
   1.1 Choose League
   1.2 Choose Prematch
2. Telco will send Message Orginated (MO) in the form of msisdn data & parameters to DCB
3. DCB sent SMS MT prediction based of Message Orginated (MO)
4. User receive SMS MT prediction
5. User Reply SMS use Code/Alias of Team
6. Telco will send Message Orginated (MO) in the form of msisdn data & others parameter to DCB
7. DCB hit ProfileAndBal API (msisdn user) to telco
   7.1 If user enough balance
   7.1.1 DBC hit DeductFee API (msisdn user) to telco (fee subscription service)
   7.1.2 DCB sent SMS MT (Sub/Welcome)
   7.1.3 User receive SMS MT (Sub/Welcome)
   7.2 If user insuff balance
   7.2.1 DCB sent SMS MT Insuff Balance
   7.2.2 User receive SMS MT (Insuff Balance)
8. If the user's prediction is correct
   8.1 DBC hit DeductFee API (msisdn user, credit_amount x number_of_goals x 1) to telco
   8.2 DCB sent SMS MT User Win
   8.3 User receive SMS MT (Win)
9. if the user's prediction is wrong
   9.1 DBC hit DeductFee API (msisdn user, credit_amount x number_of_goals x -1) to telco
   9.2 DCB sent SMS MT User Loses
   9.3 User receive SMS MT (Loses)

Note :
_Is this true, and there is no subscription renewal session?_

User Get SMS MT

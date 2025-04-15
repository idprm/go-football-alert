# Don't forget to load your .env file

export $(<.env)

_Customer Interaction Management_

1 : Live match ( Display USSD only )
2 : Flash News ( Display USSD Only )
3 : Credit Goal ( Not to live at lunch )
4 : Predict and Win ( Not to live at lunch )
5 : SMS Alerte ( Follow League and/or Team ) info sent via SMS
6 : Pronostic : Data to be uploaded into CMS to be shared via SMS , there is two segment of sub in this service

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

MT_SUCCESS_PICK
`Credit Goal: Votre participation a ete enregistree. Si Lille marque des buts lors du prochain match, vous recevrez [100][F] de credit par but. [100][F]/souscription`

MT_ALREADY_PICK
`Credit Goal: Desole, vous avez deja une equipe favorite pour le prochain match de [Ligue 1]. Vous pourrez choisir une autre equipe pour la journee suivante au [944].`

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

## USSD (#101#36#)

### Main Menu

1.  #101#36# : Main Menu

### Level 1

1.  #101#36#1 : Live Match (Confirm Message)
2.  #101#36#2 : Flash News (Confirm Message)
3.  #101#36#3 : Crédit Goal (Confirm Message)
4.  #101#36#4 : Champ. Mali (Confirm Message)
5.  #101#36#5 : Prédiction (Confirm Message)
6.  #101#36#6 : SMS Alerte (Confirm Message)
7.  #101#36#7 : Kit Foot (Confirm Message)
8.  #101#36#8 : Foot Europe (Free Access)
9.  #101#36#9 : Suiv

### Level 2

1.  #101#36#1#1 : Live Match
2.  #101#36#1#2 : Schedule
3.  #101#36#1#3 : Line Up
4.  #101#36#1#4 : Match Stats
5.  #101#36#1#5 : Display Live match
6.  #101#36#1#6 : Flash News
7.  #101#36#4#1 : Results
8.  #101#36#4#2 : Standings
9.  #101#36#4#3 : Schedule
10. #101#36#4#4 : Team
11. #101#36#4#5 : Crédit Score
12. #101#36#4#6 : Crédit Goal
13. #101#36#4#7 : SMS Alerte
14. #101#36#4#8 : SMS Alerte Equipe
15. #101#36#6#1 : Kit Foot
16. #101#36#6#2 : Europe
17. #101#36#6#3 : Afrique
18. #101#36#6#4 : SMS Alerte Equipe
19. #101#36#6#5 : Foot International
20. #101#36#7#1 : Alerte Champ. Mali + Equipe
21. #101#36#7#2 : Alerte Premier League + Equipe
22. #101#36#7#3 : Alerte La Liga + Equipe
23. #101#36#7#4 : Alerte Ligue 1 + Equipe
24. #101#36#7#5 : Alerte Serie A + Equipe
25. #101#36#7#6 : Alerte Bundesligue + Equipe
26. #101#36#8#1 : Champion League
27. #101#36#8#2 : Premier League
28. #101#36#8#3 : La Liga
29. #101#36#8#4 : Ligue 1
30. #101#36#8#5 : L. Europa
31. #101#36#8#6 : Serie A
32. #101#36#8#7 : Bundesligua
33. #101#36#8#8 : Champ Portugal
34. #101#36#8#9 : Saudi League



please confirm,
REMINDER_48H : Votre abonnement au service (Service_Name) arrive à échéance dans 48h et sera renouvelé automatiquement
SUCCESS_CHARGE : Votre abonnement au service (Service_Name) a été renouvelé avec succés




Credit Score : 
Exemple : Real vs Barca 
If i chose Real, if real win, i win the defined amount

-------------------
Credit Goal 
Exemple : Real vs Barca 
If i chose Barca and the score is 4-3 
Mean barca loose but scored 3 Goal, so i win the amount defined by goal * number of goal
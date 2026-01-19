From gate 2 several question arises ?
1. Where will I add routing ?
2. Where will I add auth ?
3. Where will I add rquest timing ?
4. Where will I add rate limiting ?

Right now, my answer is -> Everywhere.
This is messy.

So in gate 3 I will extend my existing server to support:
1. Logging middleware
2. Request timing middleware

without breaking routing, duplicating logic and touching every route manually


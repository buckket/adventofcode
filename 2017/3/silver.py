#!/usr/bin/env python3

import math

number = 289326

m = math.ceil(math.sqrt(number))
m = m if m % 2 == 1 else (m+1)

origin = ((m+1)//2, (m+1)//2)

pos = (m, m - (number - (m*m)))

if ((m*m)-(1*m)+1) <= number <= (m*m):
    pos = (m, m - ((m*m) - number))
elif ((m*m)-(2*m)+2) <= number <= ((m*m)-(1*m)):
    pos = (1, m - (((m*m)-(1*m)+1) - number))
elif ((m*m)-(3*m)+3) <= number <= ((m*m)-(2*m)+1):
    pos = ((((m*m)-(2*m)+3) - number), 1)
elif ((m*m)-(3*m)) <= number <= ((m*m)-(3*m)+2):
    pos = (m, ((m*m)-(3*m)+4) - number)

print(abs(pos[0] - origin[0]) + abs(pos[1] - origin[1]))

#1 -> 1 Ziffer
#1+2=3 -> 3*3 = 3^2 = 9 Ziffern
#1+2+2=5 = 5*5 = 5^2 = 25 Ziffern

#sqrt(x) -> nächste ungerade Zahl

#9 -> 3x3 Matrix

#1 ist immer in der Mitte


#5 4 3
#6 1 2
#7 8 9

#Pos 1 = ((3+1)/2, (3+1)/2)

#Kenne Pos Endzahl
#Kenne Anfang = (3-1)*(3-1) + 1 -> 2

#dann 9-2 

#wenn r6zahl <= 3 dann

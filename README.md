# ascii-luminosity
Prints a gradient of ascii/ansi characters from darkest to lightest or vice versa.

## Parameters
```
-file=<font_file> -mode=<Lightening|Darkening> -charset=<Ascii|Ansi>
```

## Examples

```
go run main.go -font=/usr/share/fonts/truetype/firacode/FiraCodeNerdFont-Regular.ttf -mode=Lightening -charset=Ascii

@%WM$gQ0NDR8UKqOd&BGAw9m#H6XSpbE5VZhyPaek3CFuo2j*(4JnsxITzYL}r+?f]{[i71vclt/)><\=|;^"!:~,_-'.` 
```

```
go run main.go -font=/usr/share/fonts/truetype/firacode/FiraCodeNerdFont-Regular.ttf -mode=Darkening -charset=Ascii

 `.'-_,:~!"^;|=<\>tl/)cvi{7][1?r+f}LYIzTsxJn*4(j2FoCuk3eaPyVhZ5bpE6XS#Hm9wAG&BOdqKU8RDN0Qg$MW%@
```

```
go run main.go -font=/usr/share/fonts/truetype/firacode/FiraCodeNerdFont-Regular.ttf -mode=Lightening -charset=Ansi

ØÑ@¶%ÛWMÜÙÕÄÔ$ÅÚÂgÖÃÐQÒÓÁ0NÆ©ËDÀÊRð¾ßæ8âøÉÈUqôêåKëþûãBä&dOüwGAõ¤§ÿ9ýá®Çmè¥öàéñ#¼HóùSXò6pEúb½ÏÎ£5ªºÌVZÍhPyÝÞµeça3¢kFoCu2j»ïîn4J(*x«sTIzíYìL}+?fr¿{i]1[7vc/lt)><±\|=÷;^°²³×!"¦¡~:,¬_¸¹-­'¨`.´·¯  
```

```
go run main.go -font=/usr/share/fonts/truetype/firacode/FiraCodeNerdFont-Regular.ttf -mode=Lightening -charset=Ansi

  ·¯`´.¨'­-¸¹¬_,¡:~¦!"×³²°^÷;|=\<±>t)/lcv1¿[]{7i?f+rì}LíYzTI«xs(*Jn4jïî»2uCFo3k¢eçayÝÞPµÍZhÌVºÎª5£½ÏEpúbóùòX6S¼H#ñö¥èméàá®Çý¤A9§Gwõüÿdäã&OBûëþKåqêôUÈÉâø8æß¾ðRÊÀËD0N©ÆÁÒÓQÐgÂÃÖ$ÅÚÄÙÕÔÜMÛW%¶@ÑØ
```

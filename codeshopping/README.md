# Codeshopping

## projects that have helped 

* [generic shopping cart](https://github.com/fernandez14/go-cart)
* [goddd](https://github.com/marcusolsson/goddd/blob/master/main.go)

## Frontend part

* babel for transpiling between ES6 to ES2015 for most webbrowsers

Web
*	Öppna projektets rootfolder i terminalen och kör ’npm init’ (NodeJS måste vara installerat globalt på burken)
*	Hämta alla dependencies som du vill använda. Sök på ”npm någotjagvillgöra”. Användnings- och installationsinstruktioner finns för alla moduler. 
*	CSS: Ska concateneras och minifieras. Om man vill kan man använda ett superset som LESS eller SASS
*	JS: Ska concateneras och minifieras. Om man vill kan man använda ett superset som TS. Om man använder ES2015-kod behövs en transpiler som Babel eller TS.
*	Förslagsvis används en Task Runner som t.ex. Grunt, Gulp eller Webpack för att behandla bland annat JS- och CSS-filer
*	Tester ska kunna köras på något sätt. Använd t.ex. Mocha eller Jasmine för att köra testerna. Förslagsvis körs de från terminalen genom ett npm-script som definieras i package.json. Så här kan ett sådant script se ut: 
"test": "mocha \"myWebApp/**/*.test.js\" "
*	Entrypointen till siten är genom index.html. Här länkar vi till våra JS- och CSS-filer och startar igång appen. I produktionsmiljö länkar vi till de minifierade filerna, medan vi i devmiljöer kan använda de vanliga, omodifierade filerna.
Färdiga templates för att komma igång snabbt med ett projekt (GitHub): 
*	Google Web Starter Kit, Angular Seed, Angular 2 Seed, React Redux Starter Kit
Exempel-Gist på en enkel Gulp-fil: https://gist.github.com/jesperbrannstrom/fe62e9bce79b065e83196bb8cd7edbaf
*	Kör gärna ett CSS-ramverk (t.ex. Bootstrap) för att få upp något snabbt som ser bra ut.

Continuous Deployment
I den här exempelapplikationen kommer vi att sätta upp ett enkelt Continuous Deployment-flöde. Det tänkte flödet kommer se ut ungefär så här:
1.	Vi pushar koden till master: git push origin master
2.	Gitrepot reagerar på den nya koden och meddelerar Azure (eller annan molnleverantör) om att ny kod finns på plats genom en Web Hook.
3.	Vår molntjänst tar emot meddelandet och hämtar den nya koden från vårt Gitrepo.
4.	Projektet byggs (Web / Api / Domain / Etc) 
5.	Alla tester körs. Om något test failar ska vi avbryta och logga vad som gick fel.
6.	Hantera klientresurserna (JS/CSS/Bilder/Etc) i de tasks vi satt upp i vår Task Runner (se ovan). 
7.	Alla filer deployas till miljön.
8.	Sitehosten startas om och den nya versionen av siten är uppe och rullar.





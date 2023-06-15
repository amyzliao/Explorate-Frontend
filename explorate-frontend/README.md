# explorate-frontend

### Project setup
```
cd *into the explorate-frontend folder*
npm i
```

### Development server
```
npm run serve
```
The website runs in `localhost:8080`.

### Deploy to firebase
You need a [firebase](https://console.firebase.google.com/) account. Log into firebase from terminal:
```
cd *into the explorate-frontend folder*
firebase login
```
Compile changes and deploy to firebase:
```
npm run build
firebase deploy
```
The deployed website is hosted on `https://explorate-3452a.web.app/`.

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

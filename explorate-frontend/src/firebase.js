import firebase from "firebase/compat/app";
import 'firebase/compat/firestore';

import { ref, onUnmounted } from 'vue'

const firebaseConfig = {
    apiKey: "AIzaSyDzqF3PhH4gITBcecDzAF-8X_1gIjjZrtc",
    authDomain: "explorate-3452a.firebaseapp.com",
    projectId: "explorate-3452a",
    storageBucket: "explorate-3452a.appspot.com",
    messagingSenderId: "137779142600",
    appId: "1:137779142600:web:377bd89ec041d0d5f86afd",
    measurementId: "G-J2TTKWCT5P"
  };

// firebase.initializeApp(firebaseConfig);
// const db = firebase.firestore();
const app = firebase.initializeApp(firebaseConfig);
const db = app.firestore();
const oppCollection = db.collection('opportunities')

export const createOpp = opp => {
    console.log(opp)
    let result = 201
    oppCollection.add(opp)
        .then((docRef) => {
            console.log("Document written with ID: ", docRef.id)
        })
        .catch((error) => {
            console.error("Error adding document: ", error)
            result = 300
        });
    return result
}

export const getUser = async id => {
    const opp = await oppCollection.doc(id).get()
    return opp.exists ? opp.data() : null
}

export const updateOpp = (id, opp) => {
    return oppCollection.doc(id).update(opp)
}

export const deleteOpp = id => {
    return oppCollection.doc(id).delete()
}

export const useLoadOpps = () => {
    const opps = ref([])
    const close = oppCollection.onSnapshot(snapshot => {
        opps.value = snapshot.docs.map(doc => ({ id: doc.id, ...doc.data() }))
    })
    onUnmounted(close)
    return opps
}
import firebase from "firebase/compat/app";
import 'firebase/compat/firestore';
import 'firebase/compat/auth';

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

firebase.initializeApp(firebaseConfig);
const db = firebase.firestore();

export const dbCreate = (col, val) => {
    let result = 201
    db.collection(col).add(val).then((docRef) => {
        console.log("Document written with ID: ", docRef.id)
    }).catch((error) => {
        console.error("Error adding document: ", error)
        result = 301
    });

    return result;
}

export const dbGet = async (col, id) => {
    const result = await db.collection(col).doc(id).get()
    return result.exists ? result.data() : null
}

export const dbUpdate = (col, id, val) => {
    let result = 200
    db.collection(col).doc(id).update(val).then(() => {
        console.log("Successful update")
    })
    .catch((error) => {
        console.error("Error editing document: ", error)
        result = 300
    });

    return result
}

export const dbDelete = (col, id) => {
    let result = 200
    db.collection(col).doc(id).delete()
        .then(() => {
            console.log("Successful deletion")
        })
        .catch((error) => {
            console.error("Error deleting: ", error)
            result = 300
        })
    return result
}

export const dbUseLoad = col => {
    const result = ref([])
    const close = db.collection(col).orderBy('name', 'asc').onSnapshot(snapshot => {
        result.value = snapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
        }))
    })
    onUnmounted(close)
    console.log(result)
    return result
}

export const dbGetUser = (email, password) => {
    let result = db.collection('users').where("contact", "==", email).where("password", "==", password).get()

    return result
}


export const signInWithGoogle = () => {
    const provider = new firebase.auth.GoogleAuthProvider()

    let result = firebase.auth().signInWithPopup(provider)

    return result
};

export const userID = () => {
    const user = firebase.auth().currentUser
    console.log(user.uid)
    return user ? user.uid : null
}

export const userName = () => {
    const user = firebase.auth().currentUser
    console.log(user.displayName)
    return user ? user.displayName : null
}

export const userImg = () => {
    const user = firebase.auth().currentUser
    console.log(user.photoURL)
    return user ? user.photoURL : null
}

export const userEmail = () => {
    const user = firebase.auth().currentUser
    console.log(user.email)
    return user ? user.email : null
}

export const googleSignOut = () => {
    firebase.auth().signOut()
}


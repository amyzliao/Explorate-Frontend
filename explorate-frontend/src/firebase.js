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
firebase.initializeApp(firebaseConfig);
export const db = firebase.firestore();
// const oppCollection = db.collection('opportunities')
// const usersCollection = db.collection('users')

// export const createOpp = opp => {
//     console.log(opp)
//     let result = 201
//     oppCollection.add(opp)
//         .then((docRef) => {
//             console.log("Document written with ID: ", docRef.id)
//         })
//         .catch((error) => {
//             console.error("Error adding document: ", error)
//             result = 301
//         });
//     return result
// }

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

// export const getOpp = async id => {
//     const opp = await oppCollection.doc(id).get()
//     return opp.exists ? opp.data() : null
// }

export const dbGet = async (col, id) => {
    const result = await db.collection(col).doc(id).get()
    return result.exists ? result.data() : null
}

// export const updateOpp = (id, opp) => {
//     console.log("updateOpp")
//     console.log(opp)
//     console.log(id)
//     let result = 200
//     oppCollection.doc(id).update(opp)
//         .then(() => {
//             console.log("Successful update")
//         })
//         .catch((error) => {
//             console.error("Error editing document: ", error)
//             result = 300
//         });
//     return result
// }

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

// export const deleteOpp = id => {
//     console.log(id)
//     let result = 200
//     oppCollection.doc(id).delete()
//         .then(() => {
//             console.log("Successful deletion")
//         })
//         .catch((error) => {
//             console.error("Error deleting: ", error)
//             result = 300
//         })
//     return result
// }

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

// export const useLoadOpps = () => {
//     const opps = ref([])
//     const close = oppCollection.onSnapshot(snapshot => {
//         opps.value = snapshot.docs.map(doc => ({ id: doc.id, ...doc.data() }))
//     })
//     onUnmounted(close)
//     return opps
// }

export const dbUseLoad = col => {
    const result = ref([])
    const close = db.collection(col).orderBy('name', 'asc').onSnapshot(snapshot => {
        result.value = snapshot.docs.map(doc => ({
            id: doc.id,
            ...doc.data()
        }))
    })
    onUnmounted(close)
    return result
}

export const dbGetUser = (col, email, password) => {
    // let result = ref(null)
    // db.collection(col).where("contact", "==", email).where("password", "==", password).get().then((query) => {
    //     result.value = query.docs[0].data()
    //     // console.log(query.docs[0].data())
    // }).catch((error) => {
    //     console.log("error", error)
    // })

    let result = db.collection(col).where("contact", "==", email).where("password", "==", password).get()
    // .then((query) => {
    //     query.docs.data()
    // }).catch((error) => {
    //     console.log('error', error)
    // })

    // console.log(result.then((query) => {
    //     query.data()
    // }).catch((error) => {
    //     console.log('error', error)
    // }))

    return result
}
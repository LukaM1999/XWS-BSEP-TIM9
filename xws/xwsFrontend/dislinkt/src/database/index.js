import { initializeApp } from 'firebase/app'
import { getDatabase } from 'firebase/database'
import { getFirestore } from 'firebase/firestore'
import { getStorage } from 'firebase/storage'

const config = {
  apiKey: "AIzaSyDhRvNmktJ6ukHGMoLKLv2gE2FXKzAWfBU",
  authDomain: "dislinkt-777.firebaseapp.com",
  projectId: "dislinkt-777",
  storageBucket: "dislinkt-777.appspot.com",
  messagingSenderId: "964626109848",
  appId: "1:964626109848:web:9745419223cea7146887ed"
};

initializeApp(config)

export const firestoreDb = getFirestore()
export const realtimeDb = getDatabase()
export const storage = getStorage()

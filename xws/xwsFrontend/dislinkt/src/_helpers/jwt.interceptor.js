import axios from "axios";
import {store} from "@/main";

export function jwtInterceptor(){
  axios.interceptors.request.use(async config => {
      if(config.url.includes("notifications")){
        config.headers.common['Authorization'] = `Basic ${store.getters.oneSignalToken}`;
        config.headers.common['Content-Type'] = 'application/json';
        config.headers.common['Accept'] = 'application/json';
        return config;
      }
      const token = store.getters.token;
      if (token && !isTokenExpired(token)) {
          config.headers.common['Authorization'] = token;
      } else {
        store.commit('setToken', null);
        store.commit('setUser', null);
      }
      return config;
  });
}

export function isTokenExpired(token) {
  if(!token) return true;
  const payloadBase64 = token.split('.')[1];
  const decodedJson = Buffer.from(payloadBase64, 'base64').toString();
  const decoded = JSON.parse(decodedJson)
  const exp = decoded.exp;
  return (Date.now() >= exp * 1000)
}

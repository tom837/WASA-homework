import api from './axios';
import { saveuserid } from './auth-store';


export async function login(username){
        try{
            let response = await api.post("/login", {
                name:username.name,
                });
            console.log('Response:', response.data);
            await saveuserid(response.data.ID, response.data.name)
        }catch (e){
            throw new Error(e.response.data)
        }
    };


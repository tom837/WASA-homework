import { getuserid } from './auth-store';
import axios from './axios';

export async function feed(){
    try{
        let id = await getuserid();
        console.log(id)
        let response = await axios.get("/user",{
            headers :{
                "Authorization" : id
            }
        });
        console.log('Response:', response);
        return response.data;
        // Handle successful login
    }catch (e){
        throw new Error(e.toString());
    }
};

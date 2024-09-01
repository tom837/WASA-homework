import { getuserid } from './auth-store';
import axios from './axios';

export async function like(photoid){
    try{
        let id = await getuserid();
        let response = await axios.post(`/photos/${photoid}/likes`,{},{
            headers :{
                "Authorization" : id
            }
        });
        console.log('Response:', response);
        return response.status;
    }catch (e){
        throw new Error(e.toString());
    }
};


export async function unlike(photoid){
    try{
        let id = await getuserid();
        let response = await axios.delete(`/photos/${photoid}/likes`,{
            headers :{
                "Authorization" : id
            }
        });
        console.log('Response:', response);
        return response.status;
    }catch (e){
        throw new Error(e.toString());
    }
};
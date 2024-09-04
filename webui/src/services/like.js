import { getuserid } from './auth-store';
import api from './axios';

export async function like(photoid){
    try{
        let id = await getuserid();
        let response = await api.post(`/photos/${photoid}/likes`,{},{
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
        let response = await api.delete(`/photos/${photoid}/likes`,{
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

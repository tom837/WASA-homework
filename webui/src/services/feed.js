import { getuserid } from './auth-store';
import api from './axios';

export async function feed(){
    try{
        let id = await getuserid();
        let response = await api.get("/user",{
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


export async function profile(id){
    try{
        let response = await api.get(`/user/${id}`,{
            headers:{
                "Authorization": id
            }
        });
        console.log("response", response)
        return response.data
    }catch(e){
        throw new Error(e)
    }
}


export async function deletepic(photoid){
    let id = await getuserid();
    try{
        let response = await api.delete(`/photos/${photoid}`,{
            headers:{
                "Authorization":id
            }
        });
        console.log('Response:', response);
        return response.status;
    }catch (e){
        throw new Error(e.toString());
    }
}


export async function getuserlist(){
    try{
        let response = await api.get('/users')
        console.log('response:', response)
        return response.data
    }catch(e){
        throw new Error(e.toString())
    }
}

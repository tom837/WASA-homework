import { getuserid } from './auth-store';
import api from './axios';



export async function getfollowers(){
    let id = await getuserid();
    try{
        let response = await api.get(`/followers`,{
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


export async function Follow(user){
    let id = await getuserid();
    try{
        let response = await api.post(`/user/${user}/followers`, {},
            {headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};

export async function Unfollow(user){
    let id = await getuserid();
    try{
        let response = await api.delete(`/user/${user}/followers`,
            {headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};

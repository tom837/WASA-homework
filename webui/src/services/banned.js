import { getuserid } from './auth-store';
import axios from './axios';



export async function getbanned(id){
    try{
        let response = await axios.get(`/ban`,{
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


export async function Ban(user){
    let id = await getuserid();
    try{
        let response = await axios.post(`/user/${user}/ban`, {},
            {headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};

export async function Unban(user){
    let id = await getuserid();
    try{
        let response = await axios.delete(`/user/${user}/ban`,
            {headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};

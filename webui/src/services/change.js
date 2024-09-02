import { getuserid, saveuserid } from './auth-store';
import axios from './axios';



export async function changeusername(username){
    let id = await getuserid();
    try{
        let response = await axios.put(`/user`, {
            name:username.name,
            },{headers :{
                "Authorization" : id
            }});
        console.log('Response:', response);
        await saveuserid(id,username.name);
        return response.data
    }catch (e){
        throw new Error(e.toString())
    }
};
const tokenid = 'userid';
const tokenname= "username";
const storage = sessionStorage;

export function saveuserid(authToken, username) {
    storage.setItem(tokenid, authToken);
	storage.setItem(tokenname, username)
}

export function deleteuserid(){
	const exist = sessionStorage.getItem(tokenid);
	if (exist) {
		storage.removeItem(tokenid);
		storage.removeItem(tokenname)
	}
}

export function getuserid() {
	const user = sessionStorage.getItem(tokenid);
	if (user) {
		return user
	}
	return null
}

export function getusername(){
	const user = sessionStorage.getItem(tokenname);
	if (user){
		return user
	}
	return null
}

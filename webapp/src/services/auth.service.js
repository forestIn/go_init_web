import axios from 'axios';

const API_URL = 'http://localhost:8080/api/user/';

class AuthService {
  login(user) {
    return axios
      .post(API_URL + 'login', {
        username: user.username,
        password: user.password
      })
      .then(response => {
        if (response.data.status){        
          localStorage.setItem('user', JSON.stringify(response.data.account));        
      }        

        return response.data.account;
      });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    return axios.post(API_URL + 'new', {
      username: user.username,
      email: user.email,
      password: user.password
    });
  }
}

export default new AuthService();

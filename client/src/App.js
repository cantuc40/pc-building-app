//import logo from './logo.svg';
import './App.css';
//import {gql} from 'apollo-boost';
//import { useQuery } from '@apollo/client';
import { useQuery, gql } from '@apollo/client';

const GET_USER_IDS = gql`
  {
    users {
      id
      username
      password
      email
    }

  }`;

function DisplayMotherboardID() {
  const { loading, error, data } = useQuery(GET_USER_IDS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  console.log(data);

  return data.users.map(({ id, username, password, email }) => (
    <div>
      <p>{id}</p>
      <p>{password}</p>
      <p>{username}</p>
      <p>{email}</p>
    </div>

    ));
}







function App() {
  return (
    <div>
      <DisplayMotherboardID/>
    </div>
  );
}

export default App;

import {useState,useEffect} from 'react';
import Button from "./components/Button.jsx"
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Summoner,Reload,Setup,ToggleAccept} from "../wailsjs/go/main/App.js"
function App() {

      const [data,setData] = useState({})
      const [connected,setConnect] = useState(false)
      useEffect(()=>{
          Setup().then((result)=>{
          console.log(result)
          setConnect(result)
    })
        if(connected){
          Summoner().then((result)=>{
        console.log(result)
      
        setData(result)
        }) 


    }
                   
      }
      ,[])
    if(!connected){
    return(
      <div>
        <h2>Couldn't connect to the league client</h2>
      
        <button onClick={Reload}>Retry</button>

      </div>
          )
  }
  else{
    return(
        <div>
            <h3>Ryze Auto Accept {data.gameName}</h3>
           <Button/>
        </div>
       )
  }
}


export default App

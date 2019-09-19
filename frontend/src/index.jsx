import React from "react";
import ReactDOM from "react-dom";
import 'bootstrap/dist/css/bootstrap.min.css'
import 'font-awesome/css/font-awesome.min.css'
import ChatList from './chat-list.jsx'
import Chat from './chat.jsx'

function IndexPage(props){
    return (
        <React.Fragment>
        <div className="container-fluid">
            <div className="row main-container">
                <div className="col-2">
                    <ChatList/>
                </div>
                <div className="col-10">
                    <Chat/>
                </div>
            </div>
        </div>
        </React.Fragment>
    )
}

ReactDOM.render(
    <IndexPage/>,
    document.getElementById("root")
)
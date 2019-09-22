import React from "react";
import ReactDOM from "react-dom";
import 'bootstrap/dist/css/bootstrap.min.css'
import 'font-awesome/css/font-awesome.min.css'
import ChatList from './chat-list.jsx'
import Chat from './chat.jsx'
import {NewChannelModal} from "./modals.jsx"

class IndexPage extends React.Component {
    constructor(props){
        super(props)
        this.state = {
            "wss": new WebSocket("wss://drhscexf31.execute-api.us-east-1.amazonaws.com/dev"),
            "newChannelID": "",
            "channelFilter": "",
            "channels": ["item1", "item2"]
        }
        this.state.wss.addEventListener("open", (event) => {
            console.log("connection open")
            this.state.wss.send(JSON.stringify({
                "action": "listChannels"
            }))
        })
        this.state.wss.addEventListener("message", (event) => {
            var channels = JSON.parse(event.data).message
            console.log(channels)
            this.setState({
                "channels": channels
            })
        })
    }


    sendMessage(channelID, message){
        this.state.wss.send(JSON.stringify({
            "action":"sendMessage",
            "Message": {
                "ChannelIID": channelID,
                "Message": message,
            }
        }))
    }

    // handles creating channels
    onChannelIDChange(event){
        this.setState({
            "newChannelID": event.target.value,
        })
    }

    onCreateChannel() {
        this.state.wss.send(JSON.stringify({
            "action": "createChannel",
            "ChannelID": this.state.newChannelID
        }))
    }

    // handles filtering channels
    onChangeChannelFilter(event){
        this.setState({
            "channelFilter": event.target.value,
        })
    }

    onKeyChannelFilter(event){
        console.log(this.state.channelFilter)
        if(event.key === "Enter"){
            this.state.wss.send(JSON.stringify({
                "action": "listChannels",
                "FilterPrefix": this.state.channelFilter,
            }))
            return false
        }
    }

    render(){
        return (
            <React.Fragment>
            <NewChannelModal 
                onChannelIDChange={this.onChannelIDChange.bind(this)}
                onCreateChannel={this.onCreateChannel.bind(this)}
            />
            <div className="container-fluid">
                <div className="row main-container">
                    <div className="col-2">
                        <ChatList
                            channels={this.state.channels}
                            onChangeChannelFilter={this.onChangeChannelFilter.bind(this)}
                            onKeyChannelFilter={this.onKeyChannelFilter.bind(this)}
                        />
                    </div>
                    <div className="col-10">
                        <Chat/>
                    </div>
                </div>
            </div>
            </React.Fragment>
        )
    }
}

ReactDOM.render(
    <IndexPage/>,
    document.getElementById("root")
)
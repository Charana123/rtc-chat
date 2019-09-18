import React from "react";
import ReactDOM from "react-dom";

function IndexPage(props){
    return (
        <React.Fragment>
        <div class="container-fluid">
            <div class="row">
                <div class="col-2 bg-dark overflow-auto">
                    Hello
                </div>
                <div class="col-7 bg-light">
                    World
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
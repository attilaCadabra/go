<div class="centerBlock">
	<section class="logo">
		<img src="static/img/headerlogo_b.png" alt="">
	</section>
    <form method="post" accept-charset="utf-8" action="/">	
        <div class="loginPanel">
            <section class="title" style="border-bottom: 0px;">
                <h4>ログイン</h4>
                <div style="border-top: 1px solid; border-bottom: 1px solid">
                    {{.Error}}
                </div>
            </section>
            <section class="main">
                <div class="panel-body" id="js-on" style="display: block">
                    <div class="row loginId">
                        <label class="leftlabel">ログインID</label>
                        <div class="inputWrap ">
                            <input type="text" name="user_name" id="user-name" value="">
                        </div>					
                    </div>
                    <div class="row loginId">
                        <label class="leftlabel">パスワード</label>
                        <div class="inputWrap ">
                            <input type="password" name="hash" id="hash" value="">
                        </div>					
                    </div>
                    <div class="btnArea">
                        <div class="checkrow">
                            <div class="inputWrap ">
                                <input type="hidden" name="autologin" value="0">
                                <input type="checkbox" name="autologin" value="1" id="autologin" class="autologin">
                                <label for="autologin">ログイン情報を記憶する</label>
                            </div>						
                        </div>
                        <button type="submit" class="btn loginBtn">ログイン</button>					
                    </div>
                </div>
                    
                <script type="text/javascript">
                    
                    document.getElementById("js-on").style.display ="block";
                  
                </script>
            </section>
        </div>
    </form>	
</div>
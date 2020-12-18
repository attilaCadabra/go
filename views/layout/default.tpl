<!DOCTYPE html>
<html lang="ja">

<head>
    {{.Head}}
    {{.Css}}
</head>

<body>
    <header>
        <div id="spMenu">
            <a style="color: white;"><i class="fa fa-bars"></i></a>
        </div>
        <div class="header-title" style="position:relative">
            {{.HomeTitle}}
        </div>
        <div class="btnArea menus-tool-btn inlineFlex">
           
        </div>
    </header>
    <div style="min-width: 1423px;">

        <div id="contents" class="<?php echo $contentClass; ?>">
            {{.LayoutContent}}
            <div id="loading" class="three-quarters-loader">
                <!-- CSS LOADER -->
            </div>
        </div>
    </div>
    </div>
    <footer>
    </footer>
</body>

</html>
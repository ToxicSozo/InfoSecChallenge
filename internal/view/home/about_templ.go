// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/ToxicSozo/InfoSecChallenge/internal/view/layout"

func About() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mx-auto py-4\"><!-- Галерея изображений --><div class=\"grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4\"><!-- Изображение 1 --><div><label for=\"modal-1\" class=\"cursor-pointer\"><img src=\"https://sun9-26.userapi.com/impg/CgXOV9uRxf7lFhCf6XgHYXbsW42GPx4wVL58zA/L2YFAnB-mCc.jpg?size=1280x1280&amp;quality=95&amp;sign=aabba0e09f6bc905764a352d56dc783e&amp;type=album\" class=\"w-full h-48 object-cover\" alt=\"Описание фото 1\"></label> <input type=\"checkbox\" id=\"modal-1\" class=\"modal-toggle\"><div class=\"modal\"><div class=\"modal-box\"><img src=\"https://sun9-26.userapi.com/impg/CgXOV9uRxf7lFhCf6XgHYXbsW42GPx4wVL58zA/L2YFAnB-mCc.jpg?size=1280x1280&amp;quality=95&amp;sign=aabba0e09f6bc905764a352d56dc783e&amp;type=album\" alt=\"Описание фото 1\"><p class=\"py-4\">УРА ДРУЗЬЯ ПОЗВАЛИ В ДОТУ</p><div class=\"modal-action\"><label for=\"modal-1\" class=\"btn\">Закрыть</label></div></div></div></div><!-- Изображение 2 --><div><label for=\"modal-2\" class=\"cursor-pointer\"><img src=\"https://sun9-9.userapi.com/impg/LpGMbK8scsj_1LleBuaSu4QoE70R7iEoIBj35g/euoyNCRDojQ.jpg?size=807x1080&amp;quality=95&amp;sign=04a96f8eb7c621304d85703dc6da0430&amp;type=album\" class=\"w-full h-48 object-cover\" alt=\"Описание фото 2\"></label> <input type=\"checkbox\" id=\"modal-2\" class=\"modal-toggle\"><div class=\"modal\"><div class=\"modal-box\"><img src=\"https://sun9-9.userapi.com/impg/LpGMbK8scsj_1LleBuaSu4QoE70R7iEoIBj35g/euoyNCRDojQ.jpg?size=807x1080&amp;quality=95&amp;sign=04a96f8eb7c621304d85703dc6da0430&amp;type=album\" alt=\"Описание фото 2\"><p class=\"py-4\">УРА ДРУЗЬЯ ПОЗВАЛИ В ДОТУ</p><div class=\"modal-action\"><label for=\"modal-2\" class=\"btn\">Закрыть</label></div></div></div></div><!-- Изображение 3 --><div><label for=\"modal-3\" class=\"cursor-pointer\"><img src=\"https://sun9-24.userapi.com/impg/Pz24Alz5n3zzUKI8sgIAj_KepoBRgnpDIlKQtA/KYM4qfHvc5Y.jpg?size=810x1080&amp;quality=95&amp;sign=de3410a6553ad1ef0226653b61010084&amp;type=album\" class=\"w-full h-48 object-cover\" alt=\"Описание фото 3\"></label> <input type=\"checkbox\" id=\"modal-3\" class=\"modal-toggle\"><div class=\"modal\"><div class=\"modal-box\"><img src=\"https://sun9-24.userapi.com/impg/Pz24Alz5n3zzUKI8sgIAj_KepoBRgnpDIlKQtA/KYM4qfHvc5Y.jpg?size=810x1080&amp;quality=95&amp;sign=de3410a6553ad1ef0226653b61010084&amp;type=album\" alt=\"Описание фото 3\"><p class=\"py-4\">УРА ДРУЗЬЯ ПОЗВАЛИ В ДОТУ</p><div class=\"modal-action\"><label for=\"modal-3\" class=\"btn\">Закрыть</label></div></div></div></div><!-- Изображение 4 --><div><label for=\"modal-4\" class=\"cursor-pointer\"><img src=\"https://sun9-46.userapi.com/impg/SdM2j6pl6pCXtpw0uIIO48VdSEYXX21NOCqQeA/poqsIjdVpO0.jpg?size=1080x810&amp;quality=95&amp;sign=aaf431fd9ca1397aa24dcb359ed2c972&amp;type=album\" class=\"w-full h-48 object-cover\" alt=\"Описание фото 4\"></label> <input type=\"checkbox\" id=\"modal-4\" class=\"modal-toggle\"><div class=\"modal\"><div class=\"modal-box\"><img src=\"https://sun9-46.userapi.com/impg/SdM2j6pl6pCXtpw0uIIO48VdSEYXX21NOCqQeA/poqsIjdVpO0.jpg?size=1080x810&amp;quality=95&amp;sign=aaf431fd9ca1397aa24dcb359ed2c972&amp;type=album\" alt=\"Описание фото 4\"><p class=\"py-4\">друзья не позвали в доту((</p><div class=\"modal-action\"><label for=\"modal-4\" class=\"btn\">Закрыть</label></div></div></div></div></div><!-- Описание \"Обо мне\" --><div class=\"mt-8 p-4 text-center\"><div class=\"ghost text-xl\">Обо мне</div><p>Создатель сайта: fl1tzzy (мишок:))</p><p>Тест очень сложный (нет), вам понравится.</p><p>P.S. После прохождения теста на почту вы получите грамоту о прохождении моего крутого теста 😎</p></div></div><!-- JavaScript для закрытия модального окна при клике вне его области --> <script>\n            document.addEventListener(\"click\", function (event) {\n                const modals = document.querySelectorAll(\".modal\");\n                modals.forEach(modal => {\n                    const modalBox = modal.querySelector(\".modal-box\");\n                    const modalToggle = modal.previousElementSibling; // input[type=\"checkbox\"]\n\n                    // Проверяем, был ли клик вне области модального окна\n                    if (!modalBox.contains(event.target) && !modalToggle.contains(event.target)) {\n                        modalToggle.checked = false; // Закрываем модальное окно\n                    }\n                });\n            });\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base(layout.BaseProps{Title: "About"}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

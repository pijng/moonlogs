import { Translation } from "@/shared/types";

export const ru: Translation = {
  profile: {
    label: "Профиль",
    name: "Имя",
    email: "Email",
    role: "Роль",
    tags: "Теги",
    language: "Язык",
    theme: "Тема",
  },
  log_groups: {
    label: "Группы логов",
    general: "Общие",
    form: {
      actions: {
        create: "Создать группу логов",
        edit: "Редактировать группу логов",
      },
      name: {
        label: "Название",
        hint: "Название - используется в качестве текстового идентификатора для группы. Должно быть указано латиницей, в нижнем регистре, и с подчеркиваниями в качестве разделителей",
      },
      title: {
        label: "Заголовок",
        hint: "Заголовок - используется для человекочитаемого названия группы в веб-интерфейсе. Поиск групп также будет осуществляться на основе этого параметра",
      },
      description: {
        label: "Описание",
        hint: "Описание - используется для человекочитаемого описания деталей группы в веб-интерфейсе. Поиск групп также будет осуществляться на основе этого параметра",
      },
      tag: {
        label: "Выбрать тег",
      },
      retention_days: {
        label: "Дни хранения",
        hint: "Дни хранения - количество дней, в течение которых логи будут доступны после их создания. По истечении указанного количества дней логи будут удалены. Чтобы установить бесконечный срок хранения, укажите 0 или оставьте поле пустым",
      },
      group_query_fields: {
        label: "Поля запроса группы",
        hint: "Поля запроса группы - набор полей, по которым будет осуществляться группировка логов",
        fields: {
          title: {
            label: "Заголовок",
            hint: "Заголовок - используется для человекочитаемого названия поля в веб-интерфейсе для фильтрации логов",
          },
          name: {
            label: "Название",
            hint: "Название - используется в качестве текстового идентификатора для группы. Должно быть указано латиницей, в нижнем регистре, и с подчеркиваниями в качестве разделителей",
          },
        },
      },
      kinds: {
        label: "Виды",
        hint: "Виды - набор вариантов выбора, по которым будет осуществляться группировка логов",
        fields: {
          title: {
            label: "Заголовок",
            hint: "Заголовок - используется для человекочитаемого названия вида в веб-интерфейсе для фильтрации логов",
          },
          name: {
            label: "Название",
            hint: "Название - используется в качестве текстового идентификатора для вида. Должно быть указано латиницей, в нижнем регистре, и с подчеркиваниями в качестве разделителей",
          },
        },
      },
    },
    alerts: {
      delete: "Вы уверены, что хотите удалить эту группу логов?",
    },
    filters: {
      query: {
        label: "Фильтр",
        kind: "Вид",
      },
      time: {
        label: "Время",
        from: "От",
        to: "До",
      },
      level: {
        label: "Уровень",
      },
    },
  },
  buttons: {
    create: "Создать",
    settings: "Настройки",
    save: "Сохранить",
    delete: "Удалить",
    edit: "Редактировать",
    open: "Открыть",
    log_in: "Войти",
    register: "Зарегистрироваться",
    previous: "Предыдущая",
    next: "Следующая",
    clear: "Очистить",
  },
  auth: {
    email: "Email",
    password: "Пароль",
  },
  validations: {
    required: "Обязательное поле",
    invalid_email: "Неверный email",
    passwords_dont_match: "Пароли должны совпадать",
  },
  members: {
    label: "Пользователи",
    form: {
      actions: {
        create: "Создать пользователя",
        edit: "Редактировать пользователя",
      },
      name: "Имя",
      email: "Email",
      role: "Выберите роль",
      tag: {
        label: "Выбрать теги",
        hint: "Нажмите, чтобы выбрать теги",
      },
      password: "Пароль",
      confirm_password: "Подтвердите пароль",
      revoked: "Отозван",
    },
    alerts: {
      delete: "Вы уверены, что хотите удалить этого пользователя?",
    },
  },
  tags: {
    label: "Теги",
    form: {
      actions: {
        create: "Создать тег",
        edit: "Редактировать тег",
      },
      name: {
        label: "Название",
        hint: "Название - используется для человекочитаемого названия тега в веб-интерфейсе",
      },
      view_order: {
        label: "Порядок просмотра",
        hint: "Порядок просмотра - определяет приоритетный порядок тегов. Теги с меньшими значениями будут отображаться выше в списке групп логов",
      },
    },
    alerts: {
      delete: "Вы уверены, что хотите удалить этот тег?",
    },
  },
  api_tokens: {
    label: "API токены",
    form: {
      actions: {
        create: "Создать API токен",
        edit: "Редактировать API токен",
      },
      name: {
        label: "Название",
        hint: "Название - используется для указания того, какой сервис будет использовать этот API токен. Это не влияет на функциональность токена",
      },
      revoked: "Отозван",
      creation_hint:
        "Ваш токен для интеграции API успешно создан. Убедитесь, что сохраните его надежно сейчас, так как он не будет отображаться снова по соображениям безопасности",
    },
    alerts: {
      delete: "Вы уверены, что хотите удалить этот API токен?",
    },
  },
  tables: {
    log_groups: {
      time: "Время",
      level: "Уровень",
      text: "Текст",
      request: "Запрос",
      response: "Ответ",
    },
    members: {
      email: "Email",
      name: "Имя",
      role: "Роль",
      revoked: "Отозван",
      actions: "Действия",
    },
    tags: {
      name: "Название",
      view_order: "Порядок просмотра",
      actions: "Действия",
    },
    api_tokens: {
      name: "Название",
      token: "Токен",
      revoked: "Отозван",
      actions: "Действия ",
    },
  },
  pagination: {
    first_page: "1",
  },
  components: {
    search: {
      text: "Поиск",
    },
    sidebar: {
      open: "Открыть боковую панель",
    },
  },
  miscellaneous: {
    loading: "Загрузка...",
    brand: "Moonlogs",
    blank_option: "—",
    not_found: "Запрашиваемый ресурс не найден",
    forbidden: "У вас нет доступа к этому ресурсу",
    to_home: "Перейти на главную страницу",
    empty_search_result: "Поиск не вернул никаких логов. Пожалуйста, cкорректируйте фильтры или попробуйте еще раз позже",
  },
};

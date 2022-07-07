export const steps = (i18n) => [
  {
    element: '#guide',
    popover: {
      title: i18n('driver.guideBtn'),
      description: i18n('driver.guideBtnDesc'),
      position: 'left'
    }
  },
  {
    element: '#hamburger',
    popover: {
      title: i18n('driver.hamburgerBtn'),
      description: i18n('driver.hamburgerBtnDesc'),
      position: 'bottom'
    }
  },
  {
    element: '#screenFul',
    popover: {
      title: i18n('driver.fullScreen'),
      description: i18n('driver.fullScreenDesc'),
      position: 'left'
    }
  },
  {
    element: '#avatar',
    popover: {
      title: i18n('driver.avatar'),
      description: i18n('driver.avatarDesc'),
      position: 'left'
    }
  }
]

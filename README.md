# Vi Assistant CLI 🚀

vi/vim 명령어를 빠르게 검색하고 학습할 수 있는 CLI 도구입니다!

## 🎯 주요 기능

- 🔍 **빠른 검색**: 키워드로 vi 명령어 검색
- 📖 **상세 설명**: 각 명령어의 사용법과 예제 제공
- 🎓 **학습 모드**: 단계별 튜토리얼 (초보자/중급자)
- ⭐ **즐겨찾기**: 자주 사용하는 명령어 저장
- 🌍 **다국어 지원**: 한국어/영어 지원

## 🛠 설치 방법

### 1. Go 설치 (필수)

#### Ubuntu/Debian (우선권)
```bash
# 패키지 매니저를 통한 설치 (권장)
sudo apt update
sudo apt install golang-go

# 또는 최신 버전 설치 (선택사항)
wget https://go.dev/dl/go1.24.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

#### Windows
```powershell
# Chocolatey 사용
choco install golang

# 또는 공식 사이트에서 다운로드
# https://golang.org/dl/
```

#### macOS
```bash
# Homebrew 사용
brew install go

# 또는 공식 사이트에서 다운로드
```

#### 기타 Linux 배포판
```bash
# CentOS/RHEL
sudo yum install golang

# Fedora
sudo dnf install golang
```

### 2. Go 설치 확인
```bash
go version
```

### 3. 프로젝트 빌드

```bash
# 저장소 클론
git clone <repository-url>
cd viji

# 의존성 설치
go mod tidy

# 빌드 (Ubuntu/Linux)
go build -o viji

# 실행
./viji help
```

#### Windows에서 빌드
```powershell
# PowerShell에서
go build -o viji.exe

# 실행
.\viji.exe help
```

## 🚀 사용법

### 기본 명령어

```bash
# 도움말 보기
./viji

# 명령어 검색
./viji search copy

# 명령어 설명
./viji explain :wq

# 학습 모드 시작
./viji --learn beginner

# 즐겨찾기 추가
./viji fav add :x

# 즐겨찾기 목록
./viji fav list

# 언어 설정
./viji --lang ko
```

### Ubuntu 특화 사용법

```bash
# 전역 설치 (선택사항)
sudo cp viji /usr/local/bin/
viji help  # 어디서든 실행 가능

# 권한 설정
chmod +x viji

# 시스템 경로에 추가
echo 'export PATH=$PATH:$(pwd)' >> ~/.bashrc
source ~/.bashrc
```

### 예제

```bash
# 복사 관련 명령어 검색
./viji search copy
# 결과: yy, Y, p, P 등의 명령어 표시

# 저장 명령어 설명
./viji explain :w
# 결과: "현재 파일을 저장합니다. :w filename으로 다른 이름으로 저장 가능"

# 즐겨찾기에 명령어 추가
./viji fav add :wq

# 즐겨찾기 목록 보기
./viji fav list

# 초보자 튜토리얼 시작
./viji --learn beginner

# 영어로 출력
./viji --lang en search copy
```

## 📁 프로젝트 구조

```
viji/
├── cmd/
│   └── root.go          # CLI 루트 명령어
├── internal/
│   ├── search/          # 검색 기능
│   ├── explain/         # 설명 기능
│   ├── learn/           # 학습 모드
│   ├── hint/            # 힌트 시스템
│   └── favorites/       # 즐겨찾기
├── data/
│   └── commands.json    # 명령어 데이터베이스
├── main.go              # 메인 진입점
├── viji.exe             # 빌드된 실행 파일
└── README.md
```

## 🎯 사용자 대상

- **초보 개발자**: Git, SSH에서 vi에 갇힌 경우
- **시스템 관리자**: 서버에서 vi를 매일 사용하는 경우
- **학생**: vi를 처음 배우는 경우
- **Ubuntu 사용자**: WSL2나 서버 환경에서 vi 학습

## 🔧 기술 스펙

- **언어**: Go
- **응답 시간**: 최대 1초
- **지원 OS**: Ubuntu/Linux (우선), macOS, Windows
- **라이선스**: MIT
- **오프라인 지원**: 서버 의존성 없음
- **권장 환경**: Ubuntu 20.04+, WSL2

## 🚫 제약사항

- ❌ vim 플러그인이 아님
- ❌ GUI/TUI 없음 (CLI 전용)
- ❌ vi/vim 수정 없음 (도우미 전용)
- ✅ 초보자 우선 설계
- ✅ 최소한의 외부 의존성
- ✅ Ubuntu/WSL2 최적화

## 🚀 향후 계획

- [ ] Vim 플러그인 통합
- [ ] VSCode 확장 래퍼
- [ ] 커뮤니티 기여 명령어 데이터베이스
- [ ] Ubuntu PPA 패키징
- [ ] Snap 패키지 지원

## 🐛 문제 해결

### Go 빌드 오류 시
```bash
# Go 재설치 (Ubuntu)
sudo apt remove golang-go
sudo apt autoremove
sudo apt update
sudo apt install golang-go

# 또는 최신 버전으로 설치
wget https://go.dev/dl/go1.24.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### 실행 권한 문제
```bash
chmod +x viji
```

## 📄 라이선스

MIT License - 자유롭게 사용하세요!

---

**만든이**: Vi Assistant Team  
**버전**: 1.0.0  
**최종 업데이트**: 2024년  
**권장 환경**: Ubuntu 20.04+, WSL2 